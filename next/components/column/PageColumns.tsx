import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ColumnWrapper } from "./ColumnWrapper";
import { dark1MainBg } from "../../libs/colorTheme";
import { nonNullArray } from "../../libs/nonNullArray";
import { ModalFrame } from "../modal/ModalFrame";
import { useEffect } from "react";
import { useRouter } from "next/router";

const fragmentDefinition = graphql(`
  fragment PageColumnsFragment on Page {
    columns {
      ...ColumnWrapperFragment
      name
    }
    modal {
      ...ModalFrameFragment
    }
    focusColumn
    step
  }
`);

const scrollBarStyle = css`
  // scroll bar style
  ::-webkit-scrollbar {
    width: 8px;
    height: 8px;
    background-color: #252526; /* or add it to the track */
  }
  ::-webkit-scrollbar-thumb {
    background: #37373d;
    border-radius: 8px;
  }
  ::-webkit-scrollbar-corner {
    background-color: #252526;
  }
`;

export interface PageColumnsProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const PageColumns = (props: PageColumnsProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.columns) {
    return <></>;
  }
  const columns = nonNullArray(fragment.columns);

  if (!fragment.step) {
    return <></>;
  }
  const step = fragment.step;

  const desktopColumnWidth = 768;
  const desktopWidth = desktopColumnWidth;

  const Inner = () => (
    <div
      css={css`
        // on mobile, show one column only
        @media (max-width: 768px) {
          width: 100vw;
        }
        // on desktop, show one column ony
        width: ${desktopWidth}px;
        margin: 0 auto; // centering on desktop
        height: 100svh;

        overflow-y: hidden; // let inner column handle y-axis scroll
      `}
    >
      <div
        css={css`
          display: flex;
          gap: 8px;
        `}
      >
        {columns.map((col, index) => (
          <div
            key={col.name ? col.name : index}
            css={css`
              color: ${col.name === fragment.focusColumn ? "white" : "gray"};
            `}
          >
            {col.name}
          </div>
        ))}
      </div>
      <div
        css={css`
          // flex to allow multiple columns
          display: flex;
          gap: 20px;

          // carousel container
          scroll-snap-type: x mandatory;
          scroll-behavior: smooth;
          overflow-x: auto; // buttons are the only way to scroll
          overflow-y: hidden; // let inner column handle y-axis scroll
          /* ${scrollBarStyle} */
        `}
      >
        {columns.map((col, index) => {
          return (
            <div
              id={col.name ? col.name : undefined}
              key={col.name ? col.name : index}
              css={css`
                // important to avoid column-width shrink
                flex-shrink: 0;

                // carousel scrol to stop
                scroll-snap-align: start;

                // on mobile, use full screen
                @media (max-width: 768px) {
                  width: 100vw;
                  height: 80vh;
                }

                // on desktop, use fixed width
                width: ${desktopColumnWidth}px;
                height: 80vh;

                // in-column scroll for y-axis
                overflow-y: auto;
                overflow-x: auto; // not to conflict with outer carousel scroll
                ${scrollBarStyle}
              `}
            >
              <ColumnWrapper key={index} fragment={col} step={step} />
            </div>
          );
        })}
      </div>
    </div>
  );

  return fragment.modal ? (
    <ModalFrame fragment={fragment.modal}>
      <Inner />
    </ModalFrame>
  ) : (
    <Inner />
  );
};
