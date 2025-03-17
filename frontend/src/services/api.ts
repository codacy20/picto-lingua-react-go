import axios from 'axios';

const API_BASE_URL = 'http://localhost:8080/api';

// Types based on backend models
export interface Image {
  id: string;
  url: string;
  download_url: string;
  description: string;
  width: number;
  height: number;
  created_at: string;
  photographer: string;
  photographer_url: string;
  unsplash_url: string;
  attribution_string: string;
}

export interface VocabularyItem {
  word: string;
  definition: string;
  example?: string;
}

export interface Theme {
  id: string;
  name: string;
  description?: string;
}

export interface ProgressItem {
  word: string;
  status: string; // "known", "learning", "difficult"
  time_taken_ms?: number;
  seen_count: number;
  known_count: number;
}

export interface SessionData {
  theme_id: string;
  image_id: string;
  progress: Record<string, ProgressItem>;
  session_id: string;
  started_at: string;
  last_updated: string;
}

// API functions
export const getThemes = async (): Promise<Theme[]> => {
  const response = await axios.get(`${API_BASE_URL}/themes`);
  return response.data.themes;
};

export const getImages = async (theme: string): Promise<Image[]> => {
  const response = await axios.get(`${API_BASE_URL}/images`, {
    params: { theme }
  });
  return response.data.images;
};

export const getVocabulary = async (theme: string, count: number = 10): Promise<VocabularyItem[]> => {
  const response = await axios.get(`${API_BASE_URL}/vocabulary`, {
    params: { theme, count }
  });
  return response.data.vocabulary;
};

export const saveSession = async (sessionData: Partial<SessionData>): Promise<SessionData> => {
  const response = await axios.post(`${API_BASE_URL}/session`, sessionData);
  return response.data.session;
};

export const getSession = async (sessionId: string): Promise<SessionData | null> => {
  try {
    const response = await axios.get(`${API_BASE_URL}/session`, {
      params: { session_id: sessionId }
    });
    return response.data.session;
  } catch (error) {
    return null;
  }
}; 