import { useState, useEffect } from 'react';
import { Box, Flex, Heading, Container } from '@chakra-ui/react';
import ImageGallery from './components/ImageGallery';
import GameModeSelector from './components/GameModeSelector';
import FlashcardGame from './components/FlashcardGame';
import ThemeSelector from './components/ThemeSelector';
import { getThemes, getImages, getVocabulary, Theme, Image, VocabularyItem } from './services/api';

function App() {
  // State management
  const [themes, setThemes] = useState<Theme[]>([]);
  const [selectedTheme, setSelectedTheme] = useState<Theme | null>(null);
  const [images, setImages] = useState<Image[]>([]);
  const [selectedImage, setSelectedImage] = useState<Image | null>(null);
  const [vocabulary, setVocabulary] = useState<VocabularyItem[]>([]);
  const [gameMode, setGameMode] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  // Fetch themes on component mount
  useEffect(() => {
    const fetchThemes = async () => {
      try {
        setLoading(true);
        const fetchedThemes = await getThemes();
        setThemes(fetchedThemes);
        setLoading(false);
      } catch (err) {
        setError('Failed to load themes. Please try again.');
        setLoading(false);
      }
    };

    fetchThemes();
  }, []);

  // Fetch images when a theme is selected
  useEffect(() => {
    const fetchImages = async () => {
      if (selectedTheme) {
        try {
          setLoading(true);
          const fetchedImages = await getImages(selectedTheme.id);
          setImages(fetchedImages);
          setLoading(false);
        } catch (err) {
          setError('Failed to load images. Please try again.');
          setLoading(false);
        }
      }
    };

    if (selectedTheme) {
      fetchImages();
    }
  }, [selectedTheme]);

  // Fetch vocabulary when an image and game mode are selected
  useEffect(() => {
    const fetchVocabulary = async () => {
      if (selectedTheme && gameMode === 'flashcards') {
        try {
          setLoading(true);
          const fetchedVocabulary = await getVocabulary(selectedTheme.id, 10);
          setVocabulary(fetchedVocabulary);
          setLoading(false);
        } catch (err) {
          setError('Failed to load vocabulary. Please try again.');
          setLoading(false);
        }
      }
    };

    if (selectedTheme && gameMode === 'flashcards') {
      fetchVocabulary();
    }
  }, [selectedTheme, gameMode]);

  // Handle theme selection
  const handleThemeSelect = (theme: Theme) => {
    setSelectedTheme(theme);
    setSelectedImage(null);
    setGameMode(null);
  };

  // Handle image selection
  const handleImageSelect = (image: Image) => {
    setSelectedImage(image);
    setGameMode(null);
  };

  // Handle game mode selection
  const handleGameModeSelect = (mode: string) => {
    setGameMode(mode);
  };

  // Reset to home
  const handleReset = () => {
    setSelectedTheme(null);
    setSelectedImage(null);
    setGameMode(null);
  };

  return (
    <Container maxW="container.xl" py={5}>
      <Flex direction="column" gap={5}>
        <Box textAlign="center" mb={5}>
          <Heading as="h1" size="xl">Picto Lingua</Heading>
          <Heading as="h2" size="md" mt={2}>Learn Languages Through Images</Heading>
        </Box>

        {error && (
          <Box p={4} bg="red.100" color="red.800" borderRadius="md" mb={4}>
            {error}
          </Box>
        )}

        {!selectedTheme && (
          <ThemeSelector 
            themes={themes} 
            onThemeSelect={handleThemeSelect} 
            loading={loading}
          />
        )}

        {selectedTheme && !selectedImage && (
          <ImageGallery 
            images={images} 
            onImageSelect={handleImageSelect} 
            theme={selectedTheme}
            onBack={handleReset}
            loading={loading}
          />
        )}

        {selectedTheme && selectedImage && !gameMode && (
          <GameModeSelector 
            onGameModeSelect={handleGameModeSelect} 
            onBack={() => setSelectedImage(null)}
          />
        )}

        {selectedTheme && selectedImage && gameMode === 'flashcards' && (
          <FlashcardGame 
            vocabulary={vocabulary} 
            theme={selectedTheme}
            image={selectedImage}
            onBack={() => setGameMode(null)}
            loading={loading}
          />
        )}
      </Flex>
    </Container>
  );
}

export default App;
