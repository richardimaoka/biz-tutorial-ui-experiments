import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ColumnWrapper } from "./ColumnWrapper";
import { dark1MainBg } from "../../libs/colorTheme";

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

  const list = [1, 2, 3, 4, 5, 6, 7, 8];
  const desktopColumnWidth = 768;
  const desktopColumnGap = 20;
  const desktopWidth =
    fragment.columns.length > 1
      ? /* show up to two columns */ desktopColumnWidth * 2 + desktopColumnGap
      : desktopColumnWidth;

  return (
    <div
      css={css`
        // flex to allow multiple columns
        display: flex;
        gap: 20px;

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
      `}
    >
      {list.map((item, index) => (
        <div
          key={index}
          css={css`
            // important to avoid column-width shrink
            flex-shrink: 0;

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

            ::-webkit-scrollbar-thumb {
              background: #37373d;
              border-radius: 8px;
            }
            ::-webkit-scrollbar-corner {
              background-color: #252526;
            }
          `}
        >
          <div
            css={css`
              background-color: ${dark1MainBg};
              color: white;

              width: 100%;
              height: 150%;
              display: flex;
              justify-content: center;
              align-items: center;
            `}
          >
            {item}
          </div>
        </div>
      ))}
      {/* {fragment.columns.map(
        (col, index) => col && <ColumnWrapper key={index} fragment={col} />
      )} */}
    </div>
  );
};
