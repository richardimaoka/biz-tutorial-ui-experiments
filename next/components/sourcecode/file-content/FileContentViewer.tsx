import { css } from "@emotion/react";
import { useEffect, useRef } from "react";

import Prism from "prismjs";
import "prismjs/components/prism-protobuf"; //ts 7016 error suppressed by prism-fix.d.ts
// prism stylesheet /styles/prism-xxx.css is imported from /pages/_app.tsx, as global stylesheet import is only allowed there.
// https://nextjs.org/docs/messages/css-global

interface FileContentViewerProps {
  fileContent: string;
  sourceCodeHeight: number;
}

export const FileContentViewer = ({
  fileContent,
  sourceCodeHeight,
}: FileContentViewerProps): JSX.Element => {
  const ref = useRef<HTMLElement>(null);
  useEffect(() => {
    if (ref.current) {
      Prism.highlightElement(ref.current);
    }
  }, []);
  return (
    <div
      css={css`
        height: ${sourceCodeHeight}px;
        overflow: scroll; //scroll here, not to include file name tabe in the vertical scroll
        ::-webkit-scrollbar {
          width: 5px;
          height: 5px;
          background-color: #252526; /* or add it to the track */
        }
        ::-webkit-scrollbar-thumb {
          background: #a0a0a0;
          border-radius: 5px;
        }
      `}
    >
      <pre
        css={css`
          width: fit-content;
          min-height: 100%; //expand up to the outer element
        `}
      >
        <code className="language-protobuf" ref={ref}>
          {fileContent}
        </code>
      </pre>
    </div>
  );
};
