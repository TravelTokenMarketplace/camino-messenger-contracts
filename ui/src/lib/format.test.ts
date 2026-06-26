import { describe, expect, it } from "vitest";
import { shortAddress, shortRoleName, formatAmount, formatRelativeTime } from "./format";

describe("formatRelativeTime", () => {
  const now = 1_000_000_000_000; // fixed "now" in ms
  const ago = (sec: number) => Math.floor(now / 1000) - sec;

  it("says 'just now' for very recent times", () => {
    expect(formatRelativeTime(ago(10), now)).toBe("just now");
  });

  it("formats minutes, hours and days", () => {
    expect(formatRelativeTime(ago(5 * 60), now)).toBe("5m ago");
    expect(formatRelativeTime(ago(2 * 3600), now)).toBe("2h ago");
    expect(formatRelativeTime(ago(3 * 86400), now)).toBe("3d ago");
  });

  it("handles future / clock-skewed timestamps", () => {
    expect(formatRelativeTime(ago(-60), now)).toBe("in the future");
  });
});

describe("shortAddress", () => {
  it("truncates the middle", () => {
    expect(shortAddress("0x1234567890abcdef1234567890abcdef12345678")).toBe("0x1234…5678");
  });
});

describe("shortRoleName", () => {
  it("maps the default admin role to Admin", () => {
    expect(shortRoleName("DEFAULT_ADMIN_ROLE")).toBe("Admin");
  });

  it("title-cases and drops the _ROLE suffix", () => {
    expect(shortRoleName("SERVICE_ADMIN_ROLE")).toBe("Service Admin");
    expect(shortRoleName("MESSENGER_BOT_ROLE")).toBe("Messenger Bot");
  });
});

describe("formatAmount", () => {
  it("trims long fractional precision to significant digits", () => {
    const { display, full } = formatAmount("0.00200000000063198");
    expect(display).toBe("0.002");
    expect(full).toBe("0.00200000000063198");
  });

  it("groups the integer part with thousands separators", () => {
    expect(formatAmount("1850").display).toBe("1,850");
    expect(formatAmount("1234567.5").display).toBe("1,234,567.5");
  });

  it("leaves short values unchanged", () => {
    expect(formatAmount("1.5").display).toBe("1.5");
    expect(formatAmount("0").display).toBe("0");
  });

  it("keeps significant digits for very small numbers", () => {
    expect(formatAmount("0.000001234567").display).toBe("0.000001234567");
  });
});
