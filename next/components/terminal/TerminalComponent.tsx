import { css } from "@emotion/react";
import { useEffect, useRef } from "react";
import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { TerminalNodeComponent } from "./TerminalNodeComponent";

const TerminalComponent_Fragment = graphql(`
  fragment TerminalComponent_Fragment on Terminal {
    currentDirectory
    nodes {
      ...TerminalNodeComponent_Fragment
    }
  }
`);

export interface TerminalComponentProps {
  fragment: FragmentType<typeof TerminalComponent_Fragment>;
}

export const TerminalComponent = (
  props: TerminalComponentProps
): JSX.Element => {
  const fragment = useFragment(TerminalComponent_Fragment, props.fragment);
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (ref) {
      ref.current?.scrollIntoView({ behavior: "smooth" });
    }
  });

  return (
    <div
      css={css`
        background-color: #1e1e1e;
        height: 400px;
        overflow: scroll;
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
      {fragment.nodes?.map(
        (elem, index) =>
          elem && (
            //trick to use ref for my own JSX component
            <div ref={ref} key={index} className={`${index}-element`}>
              <TerminalNodeComponent
                fragment={elem}
                isLastElement={fragment.nodes?.length === index + 1}
              />
            </div>
          )
      )}
    </div>
  );
};
