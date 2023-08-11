"use client";

import { ReactNode, useEffect, useRef } from "react";

interface TerminalScrollIntoViewoProps {
  children: ReactNode;
  doScroll: boolean;
}

export const TerminalScrollIntoView = ({
  children,
  doScroll,
}: TerminalScrollIntoViewoProps) => {
  const ref = useRef<HTMLDivElement>(null);
  const scrollIntoView = () => {
    if (ref && doScroll) {
      ref.current?.scrollIntoView({ behavior: "smooth", block: "end" });
    }
  };

  useEffect(scrollIntoView);

  return <div ref={ref}>{children}</div>;
};
