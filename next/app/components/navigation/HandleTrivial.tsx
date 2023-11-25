"use client";

import { usePathname, useRouter } from "next/navigation";
import { useEffect } from "react";

interface Props {
  isTrivial?: boolean | null;
  nextStep?: string | null;
}

export function HandleTrivial(props: Props) {
  const pathname = usePathname();
  const router = useRouter();

  useEffect(() => {
    let timeoutId: number | null;
    if (props.isTrivial && props.nextStep) {
      const path = `${pathname}?step=${props.nextStep}`;
      const setTimeoutInterval = 1000; //milli-seconds

      timeoutId = window.setTimeout(
        () => router.push(path),
        setTimeoutInterval
      );
    }

    // Don't forget to clean up
    return function cleanup() {
      if (timeoutId) {
        window.clearTimeout(timeoutId);
      }
    };
  }, [pathname, props.isTrivial, props.nextStep, router]);

  return <></>;
}
