import { editor } from "monaco-editor";
import {
  MutableRefObject,
  ReactNode,
  useCallback,
  useEffect,
  useRef,
  useState,
} from "react";
import { Root, createRoot } from "react-dom/client";
import { useEditorBoundingBox } from "./useBoundingBox";

/**
 * Custm hook to handle Editor Tooltip.
 * As described in the return types below, it exposes a bit of inner details as hook's return...
 */
export function useTooltip(
  // Shared monaco-editor's editor instance, possibly null
  editorInstance: editor.IStandaloneCodeEditor | null,

  // Prop to be passed from the parent component
  // `tooltip?` (i.e.) Optional parameter, because a React hook should handle
  // if/else *inside*, to avoid conditionals before calling the hook
  tooltip?: {
    lineNumber: number;
    children: ReactNode;
    canRender: boolean;
  }
): {
  /**
   * @return boundingBoxRef: bounding box <div> element, to be passed as <div ref={bondingBoxRef}>...</div>
   * @return resizeWindow  : callback to call when
   */
  boundingBoxRef: MutableRefObject<HTMLDivElement | null>;
  resizeWindowCallback: () => void;
} {
  // React root to set the return of React.createRoot(), and call rootRef.render()
  // See - https://react.dev/reference/react-dom/client/createRoot
  const rootRef = useRef<Root | null>(null);

  // For content widget's root HTML element
  const [contentWidgetContainer] = useState<HTMLDivElement>(() =>
    document.createElement("div")
  );

  // Tooltip has dependency on the bounding box surrounding the monaco editor
  const { boundingBoxRef, rect, resizeWindowCallback } = useEditorBoundingBox();

  const resizeContentWidget = useCallback(() => {
    contentWidgetContainer.style.width = `${rect.width}px`;
  }, [contentWidgetContainer, rect]);

  // Render content widget
  useEffect(() => {
    console.log("useTooltip useEffect", tooltip);
    if (editorInstance && tooltip?.canRender) {
      // To avoid the following warning, you need to check if createRoot() is already called.
      //   > Warning: You are calling ReactDOMClient.createRoot() on a container that has already been passed to createRoot() before. Instead, call root.render() on the existing root instead if you want to update it.
      if (!rootRef.current) {
        rootRef.current = createRoot(contentWidgetContainer);
      }
      rootRef.current.render(tooltip.children);

      const contentWidget = createWidget(
        contentWidgetContainer,
        tooltip.lineNumber
      );

      editorInstance.addContentWidget(contentWidget);
      return () => {
        editorInstance.removeContentWidget(contentWidget);
      };
    }
  }, [contentWidgetContainer, editorInstance, tooltip, tooltip?.canRender]);

  // Resize content widget upon rect change
  useEffect(() => {
    resizeContentWidget();
  }, [resizeContentWidget, rect]);

  return { boundingBoxRef, resizeWindowCallback };
}

/**
 *  Required by Monaco-Editor's editorInstance.addContentWidget() method
 *  https://microsoft.github.io/monaco-editor/docs.html#interfaces/editor.ICodeEditor.html
 */
function createWidget(element: HTMLElement, lineNumber: number) {
  return {
    getId: function () {
      return "my.content.widget";
    },
    getDomNode: function () {
      return element;
    },
    getPosition: function () {
      return {
        position: {
          lineNumber: lineNumber,
          column: 1,
        },
        preference: [
          editor.ContentWidgetPositionPreference.BELOW,
          editor.ContentWidgetPositionPreference.ABOVE,
        ],
      };
    },
  };
}
