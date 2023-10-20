"use client";

// !!!!
// Can only be used via Next.js dynamic import with ssr false option,
// due to "monaco-editor" module using browser-side `navigator` inside.
// !!!!

import { ReactNode, useCallback, useEffect, useRef, useState } from "react";
import { EditorBare } from "./EditorBare";
import { useEditorInstance } from "../useEditorInstance";
import { editor } from "monaco-editor";
import { Root, createRoot } from "react-dom/client";

interface Props {
  editorText: string;
  language: string;

  // `edits` are immediately executed by useEffect,
  // so the resulting component = editorText + edits
  edits?: editor.IIdentifiedSingleEditOperation[];
  lineHeight?: number;
  tooltip?: {
    lineNumber: number;
    children: ReactNode;
  };
}

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

interface Rect {
  width: number;
  height: number;
}

// `default` export, for easier use with Next.js dynamic import
export default function EditorInnerOnlyDynamicallyImportable(props: Props) {
  const [editorInstance, onDidMount] = useEditorInstance();

  // for content widget
  const [contentWidgetContainer] = useState<HTMLDivElement>(
    document.createElement("div")
  );
  const rootRef = useRef<Root | null>(null);
  const [rect, setRect] = useState<Rect>({ height: 0, width: 0 });

  // register event listner on window resize
  useEffect(() => {
    function handleWindowResize() {
      if (editorInstance) {
        const currentWidth = Math.max(
          editorInstance.getContentWidth(),
          editorInstance.getScrollWidth()
        );
        const currentHeight = Math.max(
          editorInstance.getContentHeight(),
          editorInstance.getScrollHeight()
        );
        setRect({ width: currentWidth, height: currentHeight });
      }
    }
    window.addEventListener("resize", handleWindowResize);
    return () => window.removeEventListener("resize", handleWindowResize);
  }, [editorInstance]);

  // update editorText
  useEffect(() => {
    if (editorInstance) {
      editorInstance.setValue(props.editorText);
    }
  }, [editorInstance, props.editorText]);

  // update language
  useEffect(() => {
    const model = editorInstance?.getModel();
    if (model) {
      editor.setModelLanguage(model, props.language);
    }
  }, [editorInstance, props.language]);

  // execute edits
  useEffect(() => {
    if (editorInstance) {
      if (props.edits) {
        // adding tooltip
        editorInstance.updateOptions({ readOnly: false });
        const result = editorInstance.executeEdits("", props.edits);
        if (!result) {
          // TODO: throw error to trigger error.tsx
          console.log("executeEdits for monaco editor failed!");
        }
        editorInstance.updateOptions({ readOnly: true });
      } else {
        //removing tooltip
        editorInstance.updateOptions({ readOnly: false });
        editorInstance.trigger("", "undo", null);
        editorInstance.updateOptions({ readOnly: true });
      }
    }
  }, [editorInstance, props.edits]);

  // add content widget
  useEffect(() => {
    const tooltip = props.tooltip;

    if (editorInstance && tooltip) {
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
  }, [contentWidgetContainer, editorInstance, props.tooltip]);

  return (
    <EditorBare
      onDidMount={onDidMount}
      onChange={() => {}}
      lineHeight={props.lineHeight}
    />
  );
}
