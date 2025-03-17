import React from 'react';
import { 
  Box, 
  Select, 
  FormControl,
  FormLabel,
} from '@chakra-ui/react';

interface LanguageSelectorProps {
  selectedLanguage: string;
  onLanguageChange: (language: string) => void;
}

const LanguageSelector: React.FC<LanguageSelectorProps> = ({ 
  selectedLanguage, 
  onLanguageChange 
}) => {
  return (
    <Box mb={4}>
      <FormControl>
        <FormLabel>Language</FormLabel>
        <Select
          value={selectedLanguage}
          onChange={(e) => onLanguageChange(e.target.value)}
          width="200px"
        >
          <option value="english">English</option>
          <option value="dutch">Dutch</option>
        </Select>
      </FormControl>
    </Box>
  );
};

export default LanguageSelector; 