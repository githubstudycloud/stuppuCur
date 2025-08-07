import { format, formatDistance, parseISO } from 'date-fns';
import { truncate, kebabCase, camelCase, startCase } from 'lodash-es';

// Date formatting
export const formatDate = (date: Date | string, formatStr = 'yyyy-MM-dd'): string => {
  const dateObj = typeof date === 'string' ? parseISO(date) : date;
  return format(dateObj, formatStr);
};

export const formatRelativeTime = (date: Date | string): string => {
  const dateObj = typeof date === 'string' ? parseISO(date) : date;
  return formatDistance(dateObj, new Date(), { addSuffix: true });
};

// Number formatting
export const formatCurrency = (
  amount: number,
  currency = 'USD',
  locale = 'en-US'
): string => {
  return new Intl.NumberFormat(locale, {
    style: 'currency',
    currency,
  }).format(amount);
};

export const formatNumber = (
  num: number,
  options: Intl.NumberFormatOptions = {}
): string => {
  return new Intl.NumberFormat('en-US', options).format(num);
};

export const formatPercentage = (num: number, decimals = 2): string => {
  return `${(num * 100).toFixed(decimals)}%`;
};

// Text formatting
export const truncateText = (text: string, length = 100): string => {
  return truncate(text, { length });
};

export const slugify = (text: string): string => {
  return kebabCase(text.toLowerCase());
};

export const toCamelCase = (text: string): string => {
  return camelCase(text);
};

export const toTitleCase = (text: string): string => {
  return startCase(text);
};

export const capitalizeFirst = (text: string): string => {
  return text.charAt(0).toUpperCase() + text.slice(1);
};

// File size formatting
export const formatFileSize = (bytes: number): string => {
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
  if (bytes === 0) return '0 Bytes';
  
  const i = Math.floor(Math.log(bytes) / Math.log(1024));
  return `${Math.round(bytes / Math.pow(1024, i) * 100) / 100} ${sizes[i]}`;
};