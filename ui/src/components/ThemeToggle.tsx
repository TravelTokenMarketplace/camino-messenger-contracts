import { Moon, Sun } from "lucide-react";
import { useTheme } from "../theme/theme";

export function ThemeToggle() {
  const { theme, toggle } = useTheme();
  return (
    <button
      type="button"
      onClick={toggle}
      aria-label="Toggle dark mode"
      title={theme === "dark" ? "Switch to light mode" : "Switch to dark mode"}
      className="rounded border border-tarmac-300 p-1.5 text-tarmac-600 hover:bg-tarmac-100 dark:border-tarmac-700 dark:text-tarmac-300 dark:hover:bg-tarmac-800"
    >
      {theme === "dark" ? <Sun className="h-4 w-4" /> : <Moon className="h-4 w-4" />}
    </button>
  );
}
