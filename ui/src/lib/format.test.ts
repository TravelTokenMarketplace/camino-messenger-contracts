import { describe, expect, it } from "vitest";
import { shortAddress, shortRoleName } from "./format";

describe("shortAddress", () => {
  it("truncates the middle", () => {
    expect(shortAddress("0x1234567890abcdef1234567890abcdef12345678")).toBe(
      "0x1234…5678",
    );
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
