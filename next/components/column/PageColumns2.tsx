import { css } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { nonNullArray } from "../../libs/nonNullArray";
import { ModalFrame } from "../modal/ModalFrame";
import { ColumnWrapper } from "./ColumnWrapper";

const fragmentDefinition = graphql(`
  fragment PageColumns2Fragment on Page {
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

export interface PageColumns2Props {
  fragment: FragmentType<typeof fragmentDefinition>;
}

export const PageColumns2 = (props: PageColumns2Props): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);

  if (!fragment.columns || fragment.columns.length === 0) {
    return <></>;
  }
  if (!fragment.step) {
    return <></>;
  }

  const focusColumn = fragment.focusColumn;
  const columns = nonNullArray(fragment.columns);
  const columnCandidate = focusColumn
    ? columns.find((col) => col.name === focusColumn)
    : columns[0];
  const column = columnCandidate ? columnCandidate : columns[0];

  const desktopColumnWidth = 768;
  const desktopWidth = desktopColumnWidth;

  const step = fragment.step;

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
        height: 80svh;

        overflow-y: hidden; // let inner column handle y-axis scroll
      `}
    >
      <div
        css={css`
          width: 100%;
          height: 30px;
        `}
      >
        <div
          css={css`
            display: flex;
            gap: 10px;
          `}
        >
          {columns.map((col, index) => (
            <div
              key={col.name ? col.name : index}
              css={css`
                color: ${col.name === focusColumn ? "white" : "gray"};
              `}
            >
              {col.name}
            </div>
          ))}
        </div>
      </div>
      <ColumnWrapper fragment={column} step={step} />
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
