import { css, keyframes } from "@emotion/react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ColumnWrapper } from "./ColumnWrapper";
import { dark1MainBg } from "../../libs/colorTheme";
import { nonNullArray } from "../../libs/nonNullArray";
import { ModalFrame } from "../modal/ModalFrame";
import { useEffect, useState } from "react";
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

interface Animating {
  kind: "animating";
  toIndex: number;
  fromIndex: number;
}

interface Static {
  kind: "static";
  index: number;
}

type AnimationState = Animating | Static;

export const PageColumns = (props: PageColumnsProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const [animation, setAnimation] = useState<AnimationState>({
    kind: "static",
    index: 0,
  });

  useEffect(() => {
    console.log("useEffect", animation);
    switch (animation.kind) {
      case "animating":
        break; // do nothing
      case "static":
        if (fragment.columns && fragment.focusColumn) {
          // find toIndex from fragment.columns
          const columns = nonNullArray(fragment.columns);
          const focusColumn = fragment.focusColumn;
          const toIndex = columns.findIndex((col) => col.name === focusColumn);

          if (toIndex > -1 && toIndex !== animation.index) {
            setAnimation({
              kind: "animating",
              toIndex: toIndex,
              fromIndex: animation.index,
            });
          }
        }
        break;
      default:
        const _exhaustiveCheck: never = animation;
        return _exhaustiveCheck;
    }
  }, [fragment.columns, fragment.focusColumn, animation]);

  if (!fragment.columns) {
    return <></>;
  }
  const columns = nonNullArray(fragment.columns);

  if (!fragment.step) {
    return <></>;
  }
  const step = fragment.step;

  const columnWidth = 768;
  const columnGap = 20;

  const kf =
    animation.kind === "animating"
      ? keyframes`
        0% {
          left: ${-(columnWidth + columnGap) * animation.fromIndex}px;
        }
        100% {
          left: ${-(columnWidth + columnGap) * animation.toIndex}px;
        }`
      : keyframes``;

  const left =
    animation.kind === "static"
      ? -(columnWidth + columnGap) * animation.index
      : 0;

  const cssAnimate = css`
    // important to avoid column-width shrink
    flex-shrink: 0;

    // carousel scrol to stop
    scroll-snap-align: start;

    // on mobile, use full screen
    @media (max-width: 768px) {
      width: 100vw;
      height: 80vh;
    }

    position: relative;
    animation: ${kf} 1s forwards;

    // on desktop, use fixed width
    width: ${columnWidth}px;
    height: 80vh;

    // in-column scroll for y-axis
    overflow-y: auto;
    overflow-x: auto; // not to conflict with outer carousel scroll
    ${scrollBarStyle}
  `;

  const cssStatic = css`
    // important to avoid column-width shrink
    flex-shrink: 0;

    // carousel scrol to stop
    scroll-snap-align: start;

    // on mobile, use full screen
    @media (max-width: 768px) {
      width: 100vw;
      height: 80vh;
    }

    position: relative;
    left: ${left}px;

    // on desktop, use fixed width
    width: ${columnWidth}px;
    height: 80vh;

    // in-column scroll for y-axis
    overflow-y: auto;
    overflow-x: auto; // not to conflict with outer carousel scroll
    ${scrollBarStyle}
  `;

  const Inner = () => (
    <div
      css={css`
        // on mobile, show one column only
        @media (max-width: ${columnWidth}px) {
          width: 100vw;
        }
        // on desktop, show one column ony
        width: ${columnWidth}px;
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
          gap: ${columnGap}px;

          // carousel container
          scroll-snap-type: x mandatory;
          scroll-behavior: smooth;
          overflow-x: hidden; // buttons are the only way to scroll
          overflow-y: hidden; // let inner column handle y-axis scroll
          /* ${scrollBarStyle} */
        `}
      >
        {columns.map((col, index) => {
          return (
            <div
              id={col.name ? col.name : undefined}
              key={col.name ? col.name : index}
              css={animation.kind === "animating" ? cssAnimate : cssStatic}
              onAnimationEnd={() => {
                if (animation.kind === "animating") {
                  setAnimation({ kind: "static", index: animation.toIndex });
                }
              }}
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
