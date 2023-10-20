"use client";

// !!!!
// Can only be used via Next.js dynamic import with ssr false option,
// due to "monaco-editor" module using browser-side `navigator` inside.
// !!!!

import { ReactNode, useEffect, useRef, useState } from "react";
import { EditorBare } from "./EditorBare";
import { useEditorInstance } from "./useEditorInstance";
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

export type EditorEditableInnerProps = Props;

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

// `default` export, for easier use with Next.js dynamic import
export default function EditorEditableOnlyDynamicallyImportable(props: Props) {
  const [editorInstance, onDidMount] = useEditorInstance();
  const [contentWidgetContainer] = useState<HTMLDivElement>(
    document.createElement("div")
  );
  const rootRef = useRef<Root | null>(null);

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
      if (!rootRef.current) {
        rootRef.current = createRoot(contentWidgetContainer);
      }

      let root = rootRef.current;
      root.render(tooltip.children);

      const contentWidget = createWidget(
        contentWidgetContainer,
        tooltip.lineNumber
      );

      editorInstance.addContentWidget(contentWidget);
      return () => {
        editorInstance.removeContentWidget(contentWidget);
        root.unmount();
      };
    }
  }, [contentWidgetContainer, editorInstance, props.tooltip]);

  return <EditorBare onDidMount={onDidMount} lineHeight={props.lineHeight} />;
}
