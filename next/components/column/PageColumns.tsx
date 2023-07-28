import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ColumnWrapper } from "./ColumnWrapper";
import { dark1MainBg } from "../../libs/colorTheme";
import { nonNullArray } from "../../libs/nonNullArray";
import { ModalFrame } from "../modal/ModalFrame";

const fragmentDefinition = graphql(`
  fragment PageColumnsFragment on Page {
    columns {
      ...ColumnWrapperFragment
      name
    }
    modal {
      ...ModalFrameFragment
    }
  }
`);

export interface ColumnWrapperProps {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const PageColumns = (props: ColumnWrapperProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.columns) {
    return <></>;
  }

  const columns = nonNullArray(fragment.columns);

  const desktopColumnWidth = 768;
  const desktopWidth = desktopColumnWidth;

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

  const Inner = () => (
    <div
      css={css`
        // flex to allow multiple columns
        display: flex;
        gap: 20px;

        // on mobile, show one column only
        @media (max-width: 768px) {
          width: 100vw;
        }
        // on desktop, show one column ony
        width: ${desktopWidth}px;
        margin: 0 auto; // centering on desktop
        height: 80svh;

        // carousel container
        scroll-snap-type: x mandatory;
        scroll-behavior: smooth;
        overflow-x: auto; // buttons are the only way to scroll
        overflow-y: hidden; // let inner column handle y-axis scroll
        /* ${scrollBarStyle} */
      `}
    >
      {columns.map((col, index) => (
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
              height: 100vh;
            }

            // on desktop, use fixed width
            width: ${desktopColumnWidth}px;
            height: 100vh;

            // in-column scroll for y-axis
            overflow-y: auto;
            overflow-x: auto; // not to conflict with outer carousel scroll
            ${scrollBarStyle}
          `}
        >
          <ColumnWrapper key={index} fragment={col} />
        </div>
      ))}
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
