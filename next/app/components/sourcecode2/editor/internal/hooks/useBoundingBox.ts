import { MutableRefObject, useEffect, useRef, useState } from "react";

export interface EditorBoundingRect {
  width: number;
  height: number;
}
/**
 * Custm hook to handle window's 'resize' event
 */
export function useEditorBoundingBox(): {
  /**
   * @return boundingBoxRef: bounding box <div> element
   * @return rect          : width and height of the bounding box <div>
   */
  boundingBoxRef: MutableRefObject<HTMLDivElement | null>;
  rect: EditorBoundingRect;
} {
  const [rect, setRect] = useState<EditorBoundingRect>({ width: 0, height: 0 });
  const boundingBoxRef = useRef<HTMLDivElement | null>(null);

  function onWindowResize() {
    if (boundingBoxRef.current) {
      setRect({
        height: boundingBoxRef.current.offsetHeight,
        width: boundingBoxRef.current.offsetWidth,
      });
    }
  }

  useEffect(() => {
    if (window) {
      window.addEventListener("resize", onWindowResize);
      return () => window.removeEventListener("resize", onWindowResize);
    }
  });

  return { boundingBoxRef, rect };
}
