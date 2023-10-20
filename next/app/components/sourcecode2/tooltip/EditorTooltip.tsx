"use client";

import { ReactNode, useEffect, useRef, useState } from "react";
import styles from "./EditorTooltipCC.module.css";
import { MarkdownDefaultStyle } from "../../markdown2/client-component/MarkdownDefaultStyle";

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
  markdownBody: string;
}

export type EditorTooltipProps = Props;

export function EditorTooltip(props: Props) {
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
      <div className={styles.tooltip}>
        <MarkdownDefaultStyle markdownBody={props.markdownBody} />
      </div>
    </div>
  );
}
