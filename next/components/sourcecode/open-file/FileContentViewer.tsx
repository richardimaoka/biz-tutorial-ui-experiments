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

// prism stylesheet /styles/prism-xxx.css is imported from /pages/_app.tsx, as global stylesheet import is only allowed there.
// https://nextjs.org/docs/messages/css-global

import { FragmentType, graphql, useFragment } from "../../../libs/gql";

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
  useEffect(() => {
    if (ref.current) {
      Prism.highlightElement(ref.current);
    }
  }, []);

  // See https://prismjs.com/#basic-usage for className="language-xxxx"
  const prismLanguage = fragment.language
    ? `language-${fragment.language}`
    : undefined;

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
        data-line={"5-10"}
        css={css`
          width: auto; //if content width < parent width, then expand up to parent width
          min-width: fit-content; //if content width > parent width, expand up to the content width
          min-height: //expand up to the outer element
            calc(100% - ${scrollBarWidth}px);
        `}
      >
        {/* See https://prismjs.com/#basic-usage for className="language-xxxx". 
            Also className={undefined} removes className attribute in React. */}
        <code className={prismLanguage} ref={ref}>
          {fragment.content}
        </code>
      </pre>
    </div>
  );
};
