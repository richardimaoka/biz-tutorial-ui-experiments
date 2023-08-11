"use client";

import { ReactNode, useEffect, useRef } from "react";

interface TerminalScrollIntoViewoProps {
  children: ReactNode;
}

export const TerminalScrollIntoView = ({
  children,
}: TerminalScrollIntoViewoProps) => {
  const ref = useRef<HTMLDivElement>(null);
  const scrollIntoView = () => {
    if (ref) {
      console.log("scrollIntoView");
      ref.current?.scrollIntoView({ behavior: "smooth", block: "end" });
    }
  };

  useEffect(scrollIntoView);

  return <div ref={ref}>{children}</div>;
};
