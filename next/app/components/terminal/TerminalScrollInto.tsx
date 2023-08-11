"use client";

import { ReactNode, useEffect, useRef } from "react";

interface TerminalScrollIntoProps {
  children: ReactNode;
}

export const TerminalScrollInto = ({ children }: TerminalScrollIntoProps) => {
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
