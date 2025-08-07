// Storage interface
interface StorageAdapter {
  getItem(key: string): string | null;
  setItem(key: string, value: string): void;
  removeItem(key: string): void;
  clear(): void;
}

// Safe storage wrapper that handles SSR and errors
class SafeStorage implements StorageAdapter {
  private storage: Storage | null;

  constructor(storageType: 'localStorage' | 'sessionStorage' = 'localStorage') {
    try {
      this.storage = typeof window !== 'undefined' ? window[storageType] : null;
    } catch {
      this.storage = null;
    }
  }

  getItem(key: string): string | null {
    try {
      return this.storage?.getItem(key) ?? null;
    } catch {
      return null;
    }
  }

  setItem(key: string, value: string): void {
    try {
      this.storage?.setItem(key, value);
    } catch {
      // Silently fail if storage is not available
    }
  }

  removeItem(key: string): void {
    try {
      this.storage?.removeItem(key);
    } catch {
      // Silently fail if storage is not available
    }
  }

  clear(): void {
    try {
      this.storage?.clear();
    } catch {
      // Silently fail if storage is not available
    }
  }
}

// Type-safe storage utilities
export class TypedStorage<T = any> {
  private storage: SafeStorage;
  private prefix: string;

  constructor(
    prefix = 'app',
    storageType: 'localStorage' | 'sessionStorage' = 'localStorage'
  ) {
    this.storage = new SafeStorage(storageType);
    this.prefix = prefix;
  }

  private getKey(key: string): string {
    return `${this.prefix}:${key}`;
  }

  get(key: string): T | null {
    try {
      const item = this.storage.getItem(this.getKey(key));
      return item ? JSON.parse(item) : null;
    } catch {
      return null;
    }
  }

  set(key: string, value: T): void {
    try {
      this.storage.setItem(this.getKey(key), JSON.stringify(value));
    } catch {
      // Silently fail
    }
  }

  remove(key: string): void {
    this.storage.removeItem(this.getKey(key));
  }

  clear(): void {
    // Only clear items with our prefix
    if (typeof window !== 'undefined' && window.localStorage) {
      const keys = Object.keys(window.localStorage);
      keys.forEach(key => {
        if (key.startsWith(`${this.prefix}:`)) {
          this.storage.removeItem(key);
        }
      });
    }
  }

  has(key: string): boolean {
    return this.get(key) !== null;
  }
}

// Pre-configured instances
export const localStorage = new TypedStorage('app', 'localStorage');
export const sessionStorage = new TypedStorage('app', 'sessionStorage');

// Utility functions
export const createStorage = (
  prefix?: string,
  type: 'localStorage' | 'sessionStorage' = 'localStorage'
) => {
  return new TypedStorage(prefix, type);
};

// Cookie utilities (for SSR-safe storage)
export const cookies = {
  get(name: string): string | null {
    if (typeof document === 'undefined') return null;
    
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    
    if (parts.length === 2) {
      return decodeURIComponent(parts.pop()?.split(';').shift() || '');
    }
    
    return null;
  },

  set(
    name: string,
    value: string,
    options: {
      expires?: Date;
      maxAge?: number;
      path?: string;
      domain?: string;
      secure?: boolean;
      sameSite?: 'strict' | 'lax' | 'none';
    } = {}
  ): void {
    if (typeof document === 'undefined') return;

    let cookieString = `${name}=${encodeURIComponent(value)}`;

    if (options.expires) {
      cookieString += `; expires=${options.expires.toUTCString()}`;
    }

    if (options.maxAge) {
      cookieString += `; max-age=${options.maxAge}`;
    }

    if (options.path) {
      cookieString += `; path=${options.path}`;
    }

    if (options.domain) {
      cookieString += `; domain=${options.domain}`;
    }

    if (options.secure) {
      cookieString += '; secure';
    }

    if (options.sameSite) {
      cookieString += `; samesite=${options.sameSite}`;
    }

    document.cookie = cookieString;
  },

  remove(name: string, path = '/'): void {
    this.set(name, '', { expires: new Date(0), path });
  },
};