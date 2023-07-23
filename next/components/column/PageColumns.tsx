import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ColumnWrapper } from "./ColumnWrapper";
import { dark1MainBg } from "../../libs/colorTheme";
import { nonNullArray } from "../../libs/nonNullArray";

const fragmentDefinition = graphql(`
  fragment PageColumnsFragment on Page {
    columns {
      ...ColumnWrapperFragment
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

  const list = [1, 2, 3, 4, 5, 6, 7, 8];
  const desktopColumnWidth = 768;
  const desktopColumnGap = 20;
  const desktopWidth =
    fragment.columns.length > 1
      ? /* show up to two columns */ desktopColumnWidth * 2 + desktopColumnGap
      : desktopColumnWidth;

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

  return (
    <div
      css={css`
        // on mobile, show one column only
        @media (max-width: 768px) {
          width: 100vw;
        }

        // on desktop, show up to two columns only
        width: ${desktopWidth}px;
        margin: 0 auto; // centering on desktop

        // carousel container
        scroll-snap-type: x mandatory;
        scroll-behavior: smooth;
        overflow-x: auto;
        overflow-y: hidden; // let inner column handle y-axis scroll
        ${scrollBarStyle}

        // flex to allow multiple columns
        display: flex;
        gap: 20px;
      `}
    >
      {columns.map((col, index) => (
        <div
          key={index}
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
            overflow-x: hidden; // not to conflict with outer carousel scroll
            ${scrollBarStyle}
          `}
        >
          <ColumnWrapper key={index} fragment={col} />
          {/* <div
                css={css`
                  background-color: ${dark1MainBg};
                  color: white;

                  width: 100%;
                  height: 150%;
                  display: flex;
                  justify-content: center;
                  align-items: center;
                `}
              /> */}
        </div>
      ))}
    </div>
  );
};
