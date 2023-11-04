"use client";

// !!!!
// Can only be used via Next.js dynamic import with ssr false option,
// due to "monaco-editor" module using browser-side `navigator` inside.
// !!!!

import { ReactNode, useCallback, useEffect, useRef, useState } from "react";
import { EditorBare } from "./EditorBare";
import { useEditorInstance } from "./hooks/useEditorInstance";
import { editor } from "monaco-editor";
import { Root, createRoot } from "react-dom/client";
import styles from "./EditorInnerOnlyDynamicallyImportable.module.css";

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
  // For editor instance
  const [editorInstance, onDidMount] = useEditorInstance();

  // For content widget
  const [contentWidgetContainer] = useState<HTMLDivElement>(
    document.createElement("div")
  );
  const boundingBoxRef = useRef<HTMLDivElement | null>(null);
  const rootRef = useRef<Root | null>(null); //React root to render
  const lineHeight = useRef<number>(0);
  function setLineHeight() {
    if (editorInstance) {
      const h = editorInstance.getOption(
        65 // somehow, NOT `enum lineHeight = 66`, but `enum lineHeight - 1` works
      );
      if (typeof h === "number") {
        lineHeight.current = h;
        console.log("line hegith =", h);
      }
    }
  }

  // React hooks starts here ---------------------------------

  // To update content widget width upon initial rendering and window resize
  const onContentAreaSizeChange = useCallback(() => {
    const bbox = boundingBoxRef.current;
    if (bbox && contentWidgetContainer) {
      console.log("onContentAreaSizeChange called, width =", bbox.offsetWidth);
      contentWidgetContainer.style.width = `${bbox.offsetWidth}px`;
    }
  }, [contentWidgetContainer]);

  // Register event listner on window resize, to set rect
  useEffect(() => {
    window.addEventListener("resize", onContentAreaSizeChange);
    return () => window.removeEventListener("resize", onContentAreaSizeChange);
  }, [editorInstance, onContentAreaSizeChange]);

  // Update editorText
  useEffect(() => {
    if (editorInstance) {
      editorInstance.setValue(props.editorText);
    }
  }, [editorInstance, props.editorText]);
  // Update language
  useEffect(() => {
    const model = editorInstance?.getModel();
    if (model) {
      editor.setModelLanguage(model, props.language);
    }
  }, [editorInstance, props.language]);

  // Execute edits
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

  // Add content widget
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

  // Rendering defined in JSX ---------------------------------

  function onChange() {
    console.log("onChange");
    onContentAreaSizeChange();
    setLineHeight();
  }

  return (
    <div className={styles.component} ref={boundingBoxRef}>
      <EditorBare
        onMount={onDidMount}
        onChange={onChange}
        lineHeight={props.lineHeight}
      />
    </div>
  );
}
