"use client";
import { ReactNode, useEffect, useState } from "react";

interface Props {
  delay?: boolean;
  children: ReactNode; // tooltip component
}

export type TerminalTooltipProps = Props;

export function TerminalTooltipDelay(props: Props) {
  const [showTooltip, setShowTooltip] = useState(false);

  useEffect(() => {
    if (props.delay) {
      setTimeout(() => {
        setShowTooltip(true);
      }, 500);
    } else {
      setShowTooltip(true);
    }
  }, [props.delay]);

  return (
    <div style={{ display: showTooltip ? "block" : "none" }}>
      {props.children}
    </div>
  );
}
