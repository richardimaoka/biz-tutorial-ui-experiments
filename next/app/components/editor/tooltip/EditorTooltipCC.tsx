"use client";

import { ReactNode, useEffect, useRef, useState } from "react";
import styles from "./EditorTooltipCC.module.css";

function pxToNumber(pxValue: string): number | null {
  return pxValue.endsWith("px") ? Number(pxValue.replace("px", "")) : null;
}

function heightIncludingMargin(element: HTMLElement): number | null {
  const style = window.getComputedStyle(element);
  const marginTop = pxToNumber(style.marginTop);
  const marginBottom = pxToNumber(style.marginBottom);
  if (marginTop && marginBottom) {
    return element.offsetHeight + marginTop + marginBottom;
  } else {
    return null;
  }
}

interface Props {
  // Since Markdown component is a server component with async rehype-react,
  // client component needs to interleave with the server component using children-passing
  children: ReactNode;
}

export type EditorTooltipProps = Props;

export function EditorTooltipCC(props: Props) {
  const ref = useRef<HTMLDivElement>(null);
  const [height, setHeight] = useState(0);

  useEffect(() => {
    if (ref.current) {
      const refHeight = heightIncludingMargin(ref.current);

      if (refHeight) {
        setHeight(refHeight);
      } else {
        //TODO: throw, and handle error in error.tsx
      }
    }
  }, [ref]);

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
