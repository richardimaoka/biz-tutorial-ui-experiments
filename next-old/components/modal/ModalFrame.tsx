import { FragmentType, graphql, useFragment } from "../../libs/gql";
import { ModalPosition } from "../../libs/gql/graphql";
import { css } from "@emotion/react";
import { ReactNode } from "react";

const fragmentDefinition = graphql(`
  fragment ModalFrameFragment on Modal {
    text
    position
  }
`);

export interface ModalFrameProps {
  fragment: FragmentType<typeof fragmentDefinition>;
  children: ReactNode;
}

const ModalBox = ({ message }: { message: string }) => (
  <div
    css={css`
      //modal box styling
      padding: 8px;
      background-color: rgba(255, 255, 255, 0.9);
      box-shadow: 0px 4px 4px 0px rgba(0, 0, 0, 0.25);
      width: fit-content;
      margin: 0 auto;
      color: black;
    `}
  >
    {message}
  </div>
);

const positionCss = (p: ModalPosition): string => {
  switch (p) {
    case "TOP":
      return "top: 20px;";
    case "CENTER":
      return "top: 50%;";
    case "BOTTOM":
      return "bottom: 20px;";
  }
};

export const ModalFrame = (props: ModalFrameProps): JSX.Element => {
  const fragment = useFragment(fragmentDefinition, props.fragment);
  const topBottomPosition = fragment.position
    ? positionCss(fragment.position)
    : positionCss("CENTER"); //default position = CENTER

  if (!fragment.text) {
    return <></>;
  }

  return (
    <div
      css={css`
        position: relative; //to contain the modal box
        width: 100%; //somehow, without this, the background image's width is shrunk... not shure why...
      `}
    >
      <div
        css={css`
          //modal box positioning
          position: absolute;
          left: 0;
          ${topBottomPosition}
          width: 100%;
          z-index: 1; /* Sit on top */
        `}
      >
        <ModalBox message={fragment.text} />
      </div>
      {props.children}
    </div>
  );
};
