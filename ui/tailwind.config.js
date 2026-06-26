/**
 * Transit-board terminal identity for the Camino Messenger console.
 *
 * Two accents carry meaning, not decoration:
 *   camino (teal)     = active / confirmed / "go"
 *   departure (amber) = pending / in-transit
 * Neutrals read as a departures board: tarmac ink on cool paper.
 */
export default {
  darkMode: "class",
  content: ["./index.html", "./src/**/*.{ts,tsx}"],
  theme: {
    extend: {
      fontFamily: {
        // Signage display face, used with restraint (titles, brand, big numerals).
        display: ['"Space Grotesk"', "ui-sans-serif", "system-ui", "sans-serif"],
        // UI / body.
        sans: ['"IBM Plex Sans"', "ui-sans-serif", "system-ui", "sans-serif"],
        // All on-chain data: addresses, hashes, amounts.
        mono: ['"IBM Plex Mono"', "ui-monospace", "monospace"],
        num: ['"IBM Plex Mono"', "ui-monospace", "monospace"],
      },
      colors: {
        // Board ink — deep blue-black, the dark surface and primary dark text.
        tarmac: {
          DEFAULT: "#0B1220",
          50: "#F2F5F7",
          100: "#E3E9EE",
          200: "#C7D2DB",
          300: "#9FB0BE",
          400: "#6B8094",
          500: "#465A6E",
          600: "#2E3F52",
          700: "#1C2A3A",
          800: "#121C29",
          900: "#0B1220",
          950: "#070C15",
        },
        // Cool off-white paper (deliberately not warm cream).
        paper: {
          DEFAULT: "#EEF2F5",
          raised: "#F7F9FB",
        },
        // Camino brand-adjacent teal — the "go" signal.
        camino: {
          DEFAULT: "#12B8A6",
          50: "#E6FAF6",
          100: "#C2F1E9",
          200: "#8AE3D6",
          300: "#4FD1BF",
          400: "#22C2AD",
          500: "#12B8A6",
          600: "#0E9385",
          700: "#0C7568",
          800: "#0B5C52",
          900: "#0A4A43",
          950: "#042C28",
        },
        // Split-flap amber — in-transit / pending.
        departure: {
          DEFAULT: "#E8923A",
          50: "#FDF3E8",
          100: "#FAE1C6",
          200: "#F4C389",
          300: "#EFA559",
          400: "#E8923A",
          500: "#D67B23",
          600: "#B5621A",
          700: "#8F4C16",
          800: "#6B3912",
          900: "#4D290D",
        },
        // Danger / revert.
        signal: {
          DEFAULT: "#E5484D",
          fg: "#B42A2F",
          dark: "#FF6166",
        },
      },
      boxShadow: {
        board: "0 1px 0 0 rgba(11,18,32,0.04), 0 1px 2px 0 rgba(11,18,32,0.06)",
      },
      keyframes: {
        flap: {
          "0%, 100%": { transform: "rotateX(0deg)", opacity: "1" },
          "45%": { transform: "rotateX(-90deg)", opacity: "0.35" },
          "55%": { transform: "rotateX(90deg)", opacity: "0.35" },
        },
        lamp: {
          "0%, 100%": { opacity: "1" },
          "50%": { opacity: "0.35" },
        },
      },
      animation: {
        flap: "flap 1.6s ease-in-out infinite",
        lamp: "lamp 1.4s ease-in-out infinite",
      },
    },
  },
  plugins: [],
};
