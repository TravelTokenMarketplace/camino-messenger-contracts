import { describe, expect, it } from "vitest";
import { render, screen } from "@testing-library/react";
import { AddressDisplay } from "./AddressDisplay";

const addr = "0x1234567890abcdef1234567890abcdef12345678";

describe("AddressDisplay", () => {
  it("shows the full address by default with a copy button", () => {
    render(<AddressDisplay address={addr} />);
    expect(screen.getByText(addr)).toBeInTheDocument();
    expect(screen.getByRole("button", { name: /copy address/i })).toBeInTheDocument();
  });

  it("truncates when asked", () => {
    render(<AddressDisplay address={addr} truncate />);
    expect(screen.getByText("0x1234…5678")).toBeInTheDocument();
  });
});
