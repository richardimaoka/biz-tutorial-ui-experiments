import {
  MutableRefObject,
  useCallback,
  useEffect,
  useRef,
  useState,
} from "react";

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
  resizeWindow: () => void;
} {
  const [rect, setRect] = useState<EditorBoundingRect>({ width: 0, height: 0 });
  const boundingBoxRef = useRef<HTMLDivElement | null>(null);

  const resizeWindow = useCallback(() => {
    if (boundingBoxRef.current) {
      setRect({
        height: boundingBoxRef.current.offsetHeight,
        width: boundingBoxRef.current.offsetWidth,
      });
    }
  }, []);

  useEffect(() => {
    if (window) {
      window.addEventListener("resize", resizeWindow);
      return () => window.removeEventListener("resize", resizeWindow);
    }
  });

  return { boundingBoxRef, rect, resizeWindow };
}
