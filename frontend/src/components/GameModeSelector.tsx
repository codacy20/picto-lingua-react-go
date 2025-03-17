import React from 'react';
import { 
  Box, 
  Stack,
  Heading, 
  Button, 
  SimpleGrid, 
  Text, 
  Flex 
} from '@chakra-ui/react';

interface GameModeSelectorProps {
  onGameModeSelect: (mode: string) => void;
  onBack: () => void;
}

const GameModeSelector: React.FC<GameModeSelectorProps> = ({ onGameModeSelect, onBack }) => {
  // Game modes (currently only flashcards is active, but structure is ready for more)
  const gameModes = [
    {
      id: 'flashcards',
      name: 'Flashcards',
      description: 'Learn vocabulary with interactive flashcards',
      isActive: true,
    },
    {
      id: 'matching',
      name: 'Matching Game',
      description: 'Match words with their meanings',
      isActive: false,
    },
    {
      id: 'quiz',
      name: 'Quiz',
      description: 'Test your knowledge with a multiple-choice quiz',
      isActive: false,
    },
  ];

  return (
    <Box>
      <Flex justifyContent="space-between" alignItems="center" mb={4}>
        <Heading as="h3" size="lg">
          Choose a Game Mode
        </Heading>
        <Button onClick={onBack} size="sm" colorScheme="gray">
          Back to Images
        </Button>
      </Flex>

      <SimpleGrid columns={{ base: 1, md: 3 }} gap={6} mt={6}>
        {gameModes.map((mode) => (
          <Box
            key={mode.id}
            p={5}
            borderRadius="md"
            bg={mode.isActive ? 'blue.50' : 'gray.100'}
            opacity={mode.isActive ? 1 : 0.7}
            border="1px solid"
            borderColor={mode.isActive ? 'blue.200' : 'gray.200'}
            cursor={mode.isActive ? 'pointer' : 'not-allowed'}
            onClick={() => mode.isActive && onGameModeSelect(mode.id)}
            transition="all 0.2s"
            _hover={{ 
              transform: mode.isActive ? 'translateY(-2px)' : 'none',
              boxShadow: mode.isActive ? 'md' : 'none'
            }}
          >
            <Stack gap={3} align="center">
              <Heading as="h4" size="md">{mode.name}</Heading>
              <Text textAlign="center">{mode.description}</Text>
              {!mode.isActive && (
                <Text fontStyle="italic" fontSize="sm" color="gray.500">
                  Coming soon!
                </Text>
              )}
            </Stack>
          </Box>
        ))}
      </SimpleGrid>
    </Box>
  );
};

export default GameModeSelector; 