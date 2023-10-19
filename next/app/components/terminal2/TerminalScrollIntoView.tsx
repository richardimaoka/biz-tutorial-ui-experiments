"use client";

import { ReactNode, useEffect, useRef } from "react";

function ExtraPadding() {
  return <div style={{ height: "50px" }} />;
}

interface Props {
  children: ReactNode;
  doScroll?: boolean;
}

export const TerminalScrollIntoView = (props: Props) => {
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (ref && props.doScroll) {
      ref.current?.scrollIntoView({ behavior: "smooth", block: "end" });
    }
  });

  return (
    <div ref={ref}>
      {props.children}
      {props.doScroll && <ExtraPadding />}
    </div>
  );
};
