import { css } from "@emotion/react";
import { useEffect, useRef } from "react";

import Prism from "prismjs";
import "prismjs/components/prism-protobuf"; //ts 7016 error suppressed by prism-fix.d.ts
// prism stylesheet /styles/prism-xxx.css is imported from /pages/_app.tsx, as global stylesheet import is only allowed there.
// https://nextjs.org/docs/messages/css-global

interface FileContentViewerProps {
  fileContent: string;
  sourceCodeHeight: number;
  prismLanguage: string;
}

export const FileContentViewer = ({
  fileContent,
  sourceCodeHeight,
  prismLanguage,
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
        height: ${sourceCodeHeight}px; // always fix the height no matter how long the content is
        overflow: scroll; //scroll within file content (not to include file name tab in vertical scroll)
        ::-webkit-scrollbar {
          width: 5px;
          height: 5px;
          background-color: #252526; /* or add it to the track */
        }
        ::-webkit-scrollbar-thumb {
          background: #37373d;
          border-radius: 5px;
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
        <code className={`language-${prismLanguage}`} ref={ref}>
          {fileContent}
        </code>
      </pre>
    </div>
  );
};
