import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ColumnWrapper } from "./ColumnWrapper";

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

  return (
    <div
      css={css`
        display: flex;
        gap: 20px;
        height: 100vh;
        @media (max-width: 768px) {
          width: 100vw;
          height: 100vh;
        }
        width: ${768 * 2}px;
        height: 100vh;

        scroll-snap-type: x mandatory;
        scroll-behavior: smooth;
        overflow-x: auto;
      `}
    >
      {list.map((item, index) => (
        <div
          key={index}
          css={css`
            background-color: white;
            color: black;
            @media (max-width: 768px) {
              width: 100vw;
              height: 100vh;
            }
            width: 768px;
            height: 100vh;
            scroll-snap-align: start;
            flex-shrink: 0;
            padding-top: 100px;
            padding-left: 100px;
          `}
        >
          {item}
        </div>
      ))}
      {/* {fragment.columns.map(
        (col, index) => col && <ColumnWrapper key={index} fragment={col} />
      )} */}
    </div>
  );
};
