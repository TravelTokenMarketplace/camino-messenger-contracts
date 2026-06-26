import { describe, expect, it } from "vitest";
import { render, screen } from "@testing-library/react";
import { InlineSentence } from "./InlineSentence";
import { shortAddress } from "../../lib/format";

const FULL = "0xbBbB000000000000000000000000000000000002";

describe("InlineSentence", () => {
  it("upgrades a shortened address into a chip with the full value and a copy button", () => {
    const short = shortAddress(FULL);
    render(<InlineSentence sentence={`Booking token #5 bought by ${short}`} args={{ buyer: FULL }} />);

    // The visible short text still reads the same.
    const chip = screen.getByText(short);
    // Hover reveals the full address via the title attribute.
    expect(chip).toHaveAttribute("title", FULL);
    // And a copy button sits beside it.
    expect(screen.getByRole("button", { name: "Copy address" })).toBeInTheDocument();
  });

  it("leaves text without a resolvable address untouched", () => {
    render(<InlineSentence sentence="Off-chain payment support enabled" args={{ supportsOffChainPayment: true }} />);

    expect(screen.getByText("Off-chain payment support enabled")).toBeInTheDocument();
    expect(screen.queryByRole("button", { name: "Copy address" })).not.toBeInTheDocument();
  });
});
