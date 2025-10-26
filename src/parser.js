// Timestamp: 2025-10-26 01:12:59

const deepClone = (obj) => {
  if (obj === null || typeof obj !== 'object') return obj;
  return JSON.parse(JSON.stringify(obj));
};

