"use client";

import { usePathname, useRouter } from "next/navigation";
import { useEffect } from "react";

interface Props {
  prevStep?: string | null;
  nextStep?: string | null;
}

export function HandleKeyEvents(props: Props) {
  const pathname = usePathname();

  const router = useRouter();

  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      if (event.code === "Space") {
        if (event.shiftKey) {
          if (props.prevStep) {
            const path = `${pathname}?step=${props.prevStep}`;
            router.push(path);
          }
        } else {
          if (props.nextStep) {
            const path = `${pathname}?step=${props.nextStep}`;
            router.push(path);
          }
        }
      }
    };

    document.addEventListener("keydown", handleKeyDown);
    // Don't forget to clean up
    return function cleanup() {
      document.removeEventListener("keydown", handleKeyDown);
    };
  }, [pathname, props.nextStep, props.prevStep, router]);

  return <></>;
}
