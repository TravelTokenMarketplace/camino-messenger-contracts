import { createContext, useContext, useEffect, useState, type ReactNode } from "react";

type Theme = "light" | "dark";

interface ThemeCtx {
  theme: Theme;
  toggle: () => void;
}

const Ctx = createContext<ThemeCtx | undefined>(undefined);

const STORAGE_KEY = "cm-ui-theme";

function systemPrefersDark(): boolean {
  try {
    return typeof window !== "undefined" && window.matchMedia?.("(prefers-color-scheme: dark)").matches;
  } catch {
    return false;
  }
}

function getStored(): Theme | null {
  try {
    const v = localStorage.getItem(STORAGE_KEY);
    return v === "light" || v === "dark" ? v : null;
  } catch {
    return null;
  }
}

export function ThemeProvider({ children }: { children: ReactNode }) {
  const [theme, setTheme] = useState<Theme>(() => getStored() ?? (systemPrefersDark() ? "dark" : "light"));

  // Apply the theme class to <html>.
  useEffect(() => {
    document.documentElement.classList.toggle("dark", theme === "dark");
  }, [theme]);

  // Follow the OS preference automatically until the user explicitly toggles.
  useEffect(() => {
    if (getStored() !== null || typeof window === "undefined" || !window.matchMedia) return;
    const mq = window.matchMedia("(prefers-color-scheme: dark)");
    const handler = (e: MediaQueryListEvent) => {
      // A toggle since mount persists a choice; once that happens, stop letting
      // OS changes override the user's explicit selection.
      if (getStored() !== null) return;
      setTheme(e.matches ? "dark" : "light");
    };
    mq.addEventListener("change", handler);
    return () => mq.removeEventListener("change", handler);
  }, []);

  function toggle() {
    setTheme((t) => {
      const next = t === "dark" ? "light" : "dark";
      try {
        localStorage.setItem(STORAGE_KEY, next);
      } catch {
        // ignore
      }
      return next;
    });
  }

  return <Ctx.Provider value={{ theme, toggle }}>{children}</Ctx.Provider>;
}

export function useTheme(): ThemeCtx {
  const ctx = useContext(Ctx);
  if (!ctx) throw new Error("useTheme must be used within ThemeProvider");
  return ctx;
}
