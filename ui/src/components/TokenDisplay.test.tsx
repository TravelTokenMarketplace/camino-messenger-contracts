import { describe, expect, it } from "vitest";
import { render, screen } from "@testing-library/react";
import { TokenDisplay } from "./TokenDisplay";

const A = "0x29F37F6adCa168B79B8d9567eab9BE3fBF21db85";

describe("TokenDisplay", () => {
  it("shows symbol headline, name, and a compact address", () => {
    render(<TokenDisplay address={A} symbol="EURe" name="Monerium EUR emoney" />);
    expect(screen.getByText("EURe")).toBeInTheDocument();
    expect(screen.getByText("Monerium EUR emoney")).toBeInTheDocument();
    expect(screen.getByText("0x29F3…db85")).toBeInTheDocument();
  });

  it("falls back to a compact address headline when symbol is missing", () => {
    render(<TokenDisplay address={A} />);
    // Both headline and the secondary line render the compact form.
    expect(screen.getAllByText("0x29F3…db85").length).toBeGreaterThanOrEqual(1);
  });

  it("still shows name when symbol is missing (metadata resolves them independently)", () => {
    render(<TokenDisplay address={A} name="Monerium EUR emoney" />);
    expect(screen.getByText("Monerium EUR emoney")).toBeInTheDocument();
    expect(screen.getAllByText("0x29F3…db85").length).toBeGreaterThanOrEqual(1);
  });
});
