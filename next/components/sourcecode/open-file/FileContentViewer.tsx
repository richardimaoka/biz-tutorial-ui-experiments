import { css } from "@emotion/react";
import { useEffect, useRef } from "react";

import Prism from "prismjs";
import "prismjs/components/prism-protobuf"; //ts 7016 error suppressed by prism-fix.d.ts
// prism stylesheet /styles/prism-xxx.css is imported from /pages/_app.tsx, as global stylesheet import is only allowed there.
// https://nextjs.org/docs/messages/css-global

import { FragmentType, graphql, useFragment } from "../../../libs/gql";

const FileContentViewer_Fragment = graphql(`
  fragment FileContentViewer_Fragment on OpenFile {
    content
    language
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

  return (
    <div
      css={css`
        height: ${props.sourceCodeHeight}px; // fix the height no matter how long the content is
        overflow: scroll; //scroll within file content (not to include file name tab in vertical scroll)
        ::-webkit-scrollbar {
          width: 8px;
          height: 8px;
          background-color: #252526; /* or add it to the track */
        }
        ::-webkit-scrollbar-thumb {
          background: #2b2b30;
          border-radius: 8px;
        }
        ::-webkit-scrollbar-thumb:horizontal {
          background: #37373d;
          border-radius: 8px;
        }
        ::-webkit-scrollbar-corner {
          background-color: #252526;
        }
      `}
    >
      <pre
        css={css`
          width: fit-content;
          min-height: 100%; //expand up to the outer element
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
