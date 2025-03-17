import React, { useState } from 'react';
import {
  Box,
  Flex,
  Button,
  Text,
  Heading,
  Stack,
  Center,
  Spinner,
  Badge,
  Image,
} from '@chakra-ui/react';
import { VocabularyItem, Theme, Image as ImageType } from '../services/api';

interface FlashcardGameProps {
  vocabulary: VocabularyItem[];
  theme: Theme;
  image: ImageType;
  onBack: () => void;
  loading: boolean;
}

const FlashcardGame: React.FC<FlashcardGameProps> = ({
  vocabulary,
  theme,
  image,
  onBack,
  loading,
}) => {
  const [currentIndex, setCurrentIndex] = useState(0);
  const [showDefinition, setShowDefinition] = useState(false);
  const [progress, setProgress] = useState<Record<string, string>>({});

  if (loading) {
    return (
      <Center py={10}>
        <Spinner size="xl" />
      </Center>
    );
  }

  if (vocabulary.length === 0) {
    return (
      <Center py={10} flexDirection="column">
        <Text mb={5}>No vocabulary items found for this theme.</Text>
        <Button onClick={onBack}>Back to Game Selection</Button>
      </Center>
    );
  }

  const currentWord = vocabulary[currentIndex];
  const totalWords = vocabulary.length;
  const progressPercentage = (currentIndex / totalWords) * 100;

  const handleNext = () => {
    setShowDefinition(false);
    if (currentIndex < vocabulary.length - 1) {
      setCurrentIndex(currentIndex + 1);
    } else {
      // Show completion or restart
      setCurrentIndex(0);
    }
  };

  const handlePrevious = () => {
    setShowDefinition(false);
    if (currentIndex > 0) {
      setCurrentIndex(currentIndex - 1);
    }
  };

  const handleCardClick = () => {
    setShowDefinition(!showDefinition);
  };

  const markWord = (status: string) => {
    setProgress({
      ...progress,
      [currentWord.word]: status,
    });
    handleNext();
  };

  return (
    <Box>
      <Flex justifyContent="space-between" alignItems="center" mb={4}>
        <Heading as="h3" size="lg">
          Flashcards: {theme.name}
        </Heading>
        <Button onClick={onBack} size="sm" colorScheme="gray">
          Back to Game Selection
        </Button>
      </Flex>

      <Flex direction={{ base: 'column', md: 'row' }} gap={6}>
        <Box 
          width={{ base: '100%', md: '40%' }}
          borderRadius="lg"
          overflow="hidden"
          boxShadow="md"
          mb={{ base: 4, md: 0 }}
        >
          <Image
            src={image.url}
            alt={image.description || theme.name}
            width="100%"
            objectFit="cover"
          />
          <Box p={3} fontSize="sm" bg="gray.50">
            <Text fontStyle="italic" mb={1} color="gray.600">
              Photo by {image.photographer} on Unsplash
            </Text>
          </Box>
        </Box>

        <Stack flex={1} gap={4} align="stretch">
          <Box 
            height="8px" 
            width="100%" 
            bg="gray.200" 
            borderRadius="full" 
            overflow="hidden"
          >
            <Box 
              height="100%" 
              width={`${progressPercentage}%`} 
              bg="blue.500" 
              transition="width 0.5s ease-in-out"
            />
          </Box>
          <Text textAlign="center">
            Card {currentIndex + 1} of {totalWords}
          </Text>

          <Box
            p={8}
            borderRadius="md"
            bg="white"
            boxShadow="md"
            onClick={handleCardClick}
            cursor="pointer"
            minHeight="200px"
            display="flex"
            flexDirection="column"
            justifyContent="center"
            alignItems="center"
            textAlign="center"
            transition="all 0.3s"
            _hover={{ boxShadow: 'lg', transform: 'translateY(-2px)' }}
            position="relative"
          >
            {!showDefinition ? (
              <Stack gap={2}>
                <Heading size="lg">{currentWord.word}</Heading>
                {currentWord.dutch_word && (
                  <Text fontSize="md" color="teal.600">
                    {currentWord.dutch_word}
                  </Text>
                )}
              </Stack>
            ) : (
              <Stack gap={3}>
                <Text fontSize="lg" fontWeight="bold">
                  {currentWord.definition}
                </Text>
                {currentWord.example && (
                  <Text fontSize="md" fontStyle="italic" color="gray.600">
                    "{currentWord.example}"
                  </Text>
                )}
                {currentWord.dutch_definition && (
                  <Text fontSize="lg" fontWeight="bold" color="teal.600" mt={4}>
                    {currentWord.dutch_definition}
                  </Text>
                )}
                {currentWord.dutch_example && (
                  <Text fontSize="md" fontStyle="italic" color="teal.600">
                    "{currentWord.dutch_example}"
                  </Text>
                )}
              </Stack>
            )}
            <Text
              position="absolute"
              bottom="2"
              right="2"
              fontSize="xs"
              color="gray.500"
            >
              Click to {showDefinition ? 'see word' : 'see definition'}
            </Text>
          </Box>

          <Flex justify="center" mt={4} gap={4} wrap="wrap">
            <Button
              onClick={() => markWord('difficult')}
              colorScheme="red"
              variant="outline"
              size="sm"
            >
              Difficult
            </Button>
            <Button
              onClick={() => markWord('learning')}
              colorScheme="yellow"
              variant="outline"
              size="sm"
            >
              Still Learning
            </Button>
            <Button
              onClick={() => markWord('known')}
              colorScheme="green"
              variant="outline"
              size="sm"
            >
              I Know This
            </Button>
          </Flex>

          <Flex justify="center" gap={4}>
            <Button
              onClick={handlePrevious}
              disabled={currentIndex === 0}
              size="sm"
            >
              Previous
            </Button>
            <Button onClick={handleNext} size="sm" colorScheme="blue">
              Next
            </Button>
          </Flex>

          <Box mt={4}>
            <Heading size="sm" mb={2}>
              Session Progress:
            </Heading>
            <Flex wrap="wrap" gap={2}>
              {Object.entries(progress).map(([word, status]) => (
                <Badge
                  key={word}
                  colorScheme={
                    status === 'known'
                      ? 'green'
                      : status === 'learning'
                      ? 'yellow'
                      : 'red'
                  }
                  p={1}
                >
                  {word}
                </Badge>
              ))}
            </Flex>
          </Box>
        </Stack>
      </Flex>
    </Box>
  );
};

export default FlashcardGame; 