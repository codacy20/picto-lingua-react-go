import React from 'react';
import { 
  Box, 
  SimpleGrid, 
  Text, 
  Heading, 
  Button, 
  Center, 
  Spinner,
} from '@chakra-ui/react';
import { Theme } from '../services/api';

interface ThemeSelectorProps {
  themes: Theme[];
  onThemeSelect: (theme: Theme) => void;
  loading: boolean;
}

const ThemeSelector: React.FC<ThemeSelectorProps> = ({ themes, onThemeSelect, loading }) => {
  const bgColor = 'gray.100';
  const hoverBgColor = 'blue.100';

  if (loading) {
    return (
      <Center py={10}>
        <Spinner size="xl" />
      </Center>
    );
  }

  return (
    <Box>
      <Heading as="h3" size="lg" mb={4} textAlign="center">Select a Theme</Heading>
      <SimpleGrid columns={{ base: 1, md: 2, lg: 3 }} gap={6}>
        {themes.map((theme) => (
          <Box 
            key={theme.id}
            p={5}
            borderRadius="md"
            bg={bgColor}
            _hover={{ bg: hoverBgColor, transform: 'translateY(-2px)' }}
            transition="all 0.2s"
            cursor="pointer"
            onClick={() => onThemeSelect(theme)}
            boxShadow="md"
          >
            <Heading as="h4" size="md" mb={2}>{theme.name}</Heading>
            {theme.description && (
              <Text color="gray.600">{theme.description}</Text>
            )}
          </Box>
        ))}
      </SimpleGrid>
    </Box>
  );
};

export default ThemeSelector; 