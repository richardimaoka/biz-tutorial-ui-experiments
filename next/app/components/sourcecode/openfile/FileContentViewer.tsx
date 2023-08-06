"use client";

import { useEffect, useRef } from "react";

import Prism from "prismjs";

import "./prism-vsc-dark-plus.css";
import "./prism-line-numbers.css";
import "./prism-line-highlight.css";

// Prism.js plugins
// Side-effect only import - https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import#import_a_module_for_its_side_effects_only
import "prismjs/plugins/line-numbers/prism-line-numbers"; //ts 7016 error suppressed by prism-fix.d.ts in /libs folder
import "prismjs/plugins/line-highlight/prism-line-highlight"; //ts 7016 error suppressed by prism-fix.d.ts in /libs folder

// Prism.js language supports.
// Side-effect only import - https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import#import_a_module_for_its_side_effects_only
import "prismjs/components/prism-protobuf"; //ts 7016 error suppressed by prism-fix.d.ts in /libs folder
import "prismjs/components/prism-json"; //ts 7016 error suppressed by prism-fix.d.ts in /libs folder
import "prismjs/components/prism-typescript"; //ts 7016 error suppressed by prism-fix.d.ts in /libs folder

// prism stylesheet /styles/prism-xxx.css is imported from /pages/_app.tsx, as global stylesheet import is only allowed there.
// https://nextjs.org/docs/messages/css-global

import { FragmentType, graphql, useFragment } from "@/libs/gql";
import { nonNullArray } from "@/libs/nonNullArray";
import styles from "./style.module.css";

const fragmentDefinition = graphql(`
  fragment FileContentViewer_Fragment on OpenFile {
    content
    language
    highlight {
      fromLine
      toLine
    }
  }
`);

export interface FileContentViewerProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const FileContentViewer = (
  props: FileContentViewerProps
): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const ref = useRef<HTMLElement>(null);

  // See https://prismjs.com/#basic-usage for className="language-xxxx"
  const prismLanguage = fragment.language
    ? `language-${fragment.language}`
    : undefined;

  // https://prismjs.com/plugins/line-highlight/
  const dataLine = fragment.highlight
    ? nonNullArray(fragment.highlight) // remove nulls to simplify type handling in filter() and map()
        .filter((h) => h.fromLine && h.toLine) // if both is null, don't highlight it
        .map(
          // two possible styles of data-line attribute: `5` (single-line) and `1-2` (multi-line)
          (h) =>
            h.fromLine === h.toLine
              ? `${h.fromLine}` //(e.g.) 5
              : `${h.fromLine}-${h.toLine}` //(e.g.) 1-2
        )
        .join(", ") //1-2, 5, 9-20
    : "";

  useEffect(() => {
    if (ref.current) {
      // 1. Walkaround - need to set className here, not in JSX.
      //    Otherwise, a warning like below will be generated:
      //      Warning: Prop `className` did not match. Server: "line-numbers css-1rw6e6m-FileContentViewer language-json" Client: "line-numbers css-1rw6e6m-FileContentViewer"
      //        at pre
      //        See more info here: https://nextjs.org/docs/messages/react-hydration-error
      //
      // 2. Why set className="language-xxxx" to <code> (`ref` points to <code>), not <pre>? See https://prismjs.com/#basic-usage
      ref.current.className = prismLanguage ? prismLanguage : "";

      Prism.highlightElement(ref.current);
    }
  }, [fragment, prismLanguage]);

  return (
    <div className={styles.contents}>
      <pre className="line-numbers" data-line={dataLine}>
        <code ref={ref}>{fragment.content}</code>
      </pre>
    </div>
  );
};
