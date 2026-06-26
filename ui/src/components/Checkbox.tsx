import { type ReactNode } from "react";
import { Checkbox as HCheckbox, Field, Label } from "@headlessui/react";
import { Check } from "lucide-react";

interface CheckboxProps {
  checked: boolean;
  onChange: (checked: boolean) => void;
  disabled?: boolean;
  label?: ReactNode;
}

/** Styled checkbox built on HeadlessUI, with an optional clickable label. */
export function Checkbox({ checked, onChange, disabled, label }: CheckboxProps) {
  return (
    <Field
      disabled={disabled}
      className="flex items-center gap-2 text-sm data-[disabled]:cursor-not-allowed data-[disabled]:opacity-50"
    >
      <HCheckbox
        checked={checked}
        onChange={onChange}
        className="group flex h-4 w-4 shrink-0 items-center justify-center rounded border border-tarmac-300 bg-paper-raised transition-colors focus:outline-none focus-visible:ring-2 focus-visible:ring-camino-500 data-[checked]:border-camino-600 data-[checked]:bg-camino-600 dark:border-tarmac-600 dark:bg-tarmac-900 dark:data-[checked]:border-camino-500 dark:data-[checked]:bg-camino-500"
      >
        <Check className="h-3 w-3 text-white opacity-0 transition-opacity group-data-[checked]:opacity-100" />
      </HCheckbox>
      {label && <Label className="cursor-pointer select-none">{label}</Label>}
    </Field>
  );
}
