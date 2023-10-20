"use client";

import { ReactNode, useEffect, useRef, useState } from "react";
import styles from "./EditorTooltipCC.module.css";

function pxToNumber(pxValue: string): number | null {
  return pxValue.endsWith("px") ? Number(pxValue.replace("px", "")) : null;
}

interface Props {
  children: ReactNode;
}

export type EditorTooltipProps = Props;

export function EditorTooltipCC(props: Props) {
  const ref = useRef<HTMLDivElement>(null);
  const [height, setHeight] = useState(0);

  useEffect(() => {
    if (ref.current) {
      const style = window.getComputedStyle(ref.current);
      const marginTop = pxToNumber(style.marginTop);
      const marginBottom = pxToNumber(style.marginBottom);
      if (marginTop && marginBottom) {
        setHeight(ref.current.offsetHeight + marginTop + marginBottom);
      } else {
        //TODO: throw, and handle error in error.tsx
      }
    }
  }, []);

  useEffect(() => {
    if (ref.current) {
      console.log("tooltip height = " + height);
    }
  }, [height]);

  return (
    <div ref={ref} className={styles.component}>
      <div className={styles.tooltip}>{props.children}</div>
    </div>
  );
}
