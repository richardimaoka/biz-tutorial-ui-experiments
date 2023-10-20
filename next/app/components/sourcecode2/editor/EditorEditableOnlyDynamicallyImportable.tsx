"use client";

// !!!!
// Can only be used via Next.js dynamic import with ssr false option,
// due to "monaco-editor" module using browser-side `navigator` inside.
// !!!!

import { useEffect } from "react";
import { EditorBare } from "./EditorBare";
import { useEditorInstance } from "./useEditorInstance";
import { editor } from "monaco-editor";

interface Props {
  editorText: string;
  language: string;

  // `edits` are immediately executed by useEffect,
  // so the resulting component = editorText + edits
  edits?: editor.IIdentifiedSingleEditOperation[];
  lineHeight?: number;
  contentWidgetElement?: HTMLElement;
}

export type EditorEditableInnerProps = Props;

// `default` export, for easier use with Next.js dynamic import
export default function EditorEditableOnlyDynamicallyImportable(props: Props) {
  const [editorInstance, onDidMount] = useEditorInstance();

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
    const domNode = props.contentWidgetElement;
    if (editorInstance && domNode) {
      const contentWidget = {
        getId: function () {
          return "my.content.widget";
        },
        getDomNode: function () {
          return domNode;
        },
        getPosition: function () {
          return {
            position: {
              lineNumber: 7,
              column: 8,
            },
            preference: [
              editor.ContentWidgetPositionPreference.ABOVE,
              editor.ContentWidgetPositionPreference.BELOW,
            ],
          };
        },
      };
      editorInstance.addContentWidget(contentWidget);
    }
  }, [editorInstance, props.contentWidgetElement]);

  return <EditorBare onDidMount={onDidMount} lineHeight={props.lineHeight} />;
}
