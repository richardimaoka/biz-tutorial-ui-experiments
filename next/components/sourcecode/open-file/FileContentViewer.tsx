import { css } from "@emotion/react";
import { useEffect, useRef } from "react";

import Prism from "prismjs";

// Prism.js plugins
// Side-effect only import - https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import#import_a_module_for_its_side_effects_only
import "prismjs/plugins/line-numbers/prism-line-numbers"; //ts 7016 error suppressed by prism-fix.d.ts in /libs folder
import "prismjs/plugins/line-highlight/prism-line-highlight"; //ts 7016 error suppressed by prism-fix.d.ts in /libs folder

// Prism.js language supports.
// Side-effect only import - https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Statements/import#import_a_module_for_its_side_effects_only
import "prismjs/components/prism-protobuf"; //ts 7016 error suppressed by prism-fix.d.ts in /libs folder
import "prismjs/components/prism-json"; //ts 7016 error suppressed by prism-fix.d.ts in /libs folder

// prism stylesheet /styles/prism-xxx.css is imported from /pages/_app.tsx, as global stylesheet import is only allowed there.
// https://nextjs.org/docs/messages/css-global

import { FragmentType, graphql, useFragment } from "../../../libs/gql";
import { nonNullArray } from "../../../libs/nonNullArray";

const FileContentViewer_Fragment = graphql(`
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
  fragment: FragmentType<typeof FileContentViewer_Fragment>;
  sourceCodeHeight: number;
}

export const FileContentViewer = (
  props: FileContentViewerProps
): JSX.Element => {
  const fragment = useFragment(FileContentViewer_Fragment, props.fragment);
  const ref = useRef<HTMLElement>(null);

  // See https://prismjs.com/#basic-usage for className="language-xxxx"
  const prismLanguage = fragment.language
    ? `language-${fragment.language}`
    : undefined;

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

  const scrollBarWidth = 8; //px

  return (
    <div
      css={css`
        height: ${props.sourceCodeHeight}px; // fix the height no matter how long the content is
        overflow: auto; //scroll within file content (not to include file name tab in vertical scroll)
        ::-webkit-scrollbar {
          width: 8px;
          height: 8px;
          background-color: #252526; /* or add it to the track */
        }
        ::-webkit-scrollbar-thumb {
          background-color: #2b2b30;
          border-radius: ${scrollBarWidth}px;
        }
        ::-webkit-scrollbar-thumb:horizontal {
          background-color: #37373d;
          border-radius: ${scrollBarWidth}px;
        }
        ::-webkit-scrollbar-corner {
          background-color: #252526;
        }
      `}
    >
      <pre
        className="line-numbers"
        data-line={dataLine}
        css={css`
          width: auto; //if content width < parent width, then expand up to parent width
          min-width: fit-content; //if content width > parent width, expand up to the content width
          min-height: //expand up to the outer element
            calc(100% - ${scrollBarWidth}px);
        `}
      >
        <code ref={ref}>{fragment.content}</code>
      </pre>
    </div>
  );
};
