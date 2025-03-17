import React from 'react';
import { 
  Box, 
  SimpleGrid, 
  Image as ChakraImage, 
  Text, 
  Heading, 
  Button, 
  Center, 
  Spinner, 
  Link, 
  Flex
} from '@chakra-ui/react';
import { Image, Theme } from '../services/api';

interface ImageGalleryProps {
  images: Image[];
  theme: Theme;
  onImageSelect: (image: Image) => void;
  onBack: () => void;
  loading: boolean;
}

const ImageGallery: React.FC<ImageGalleryProps> = ({ 
  images, 
  theme, 
  onImageSelect, 
  onBack, 
  loading 
}) => {
  if (loading) {
    return (
      <Center py={10}>
        <Spinner size="xl" />
      </Center>
    );
  }

  return (
    <Box>
      <Flex justifyContent="space-between" alignItems="center" mb={4}>
        <Heading as="h3" size="lg">
          Images for "{theme.name}"
        </Heading>
        <Button onClick={onBack} size="sm" colorScheme="gray">
          Back to Themes
        </Button>
      </Flex>

      {images.length === 0 ? (
        <Center py={10}>
          <Text>No images found for this theme. Please try another theme.</Text>
        </Center>
      ) : (
        <SimpleGrid columns={{ base: 1, sm: 2, md: 3 }} gap={6}>
          {images.map((image) => (
            <Box 
              key={image.id}
              borderRadius="lg"
              overflow="hidden"
              boxShadow="md"
              transition="all 0.3s"
              _hover={{ transform: 'scale(1.02)', boxShadow: 'lg' }}
              cursor="pointer"
              onClick={() => onImageSelect(image)}
            >
              <Box position="relative">
                <ChakraImage 
                  src={image.url} 
                  alt={image.description || theme.name}
                  width="100%"
                  height="200px"
                  objectFit="cover"
                />
              </Box>
              
              <Box p={3} fontSize="sm" bg="gray.50">
                <Text fontStyle="italic" mb={1} color="gray.600">
                  Photo by{' '}
                  <Link 
                    href={image.photographer_url} 
                    color="blue.500"
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    {image.photographer}
                  </Link>
                  {' '}on{' '}
                  <Link 
                    href={image.unsplash_url} 
                    color="blue.500"
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    Unsplash
                  </Link>
                </Text>
                {image.description && (
                  <Text fontSize="sm" overflow="hidden" height="40px">{image.description}</Text>
                )}
              </Box>
            </Box>
          ))}
        </SimpleGrid>
      )}
    </Box>
  );
};

export default ImageGallery; 