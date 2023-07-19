import { ReactNode } from "react";
import { ImageDescriptionColumnPosition } from "../../libs/gql/graphql";
import { dark1MainBg } from "../../libs/colorTheme";
import { css } from "@emotion/react";

interface ColumnContentsPositionProps {
  position: ImageDescriptionColumnPosition;
  children: ReactNode;
}

export const ColumnContentsPosition = ({
  position,
  children,
}: ColumnContentsPositionProps) => {
  const justifyContent = (p: ImageDescriptionColumnPosition): string => {
    switch (p) {
      case "TOP":
        return "flex-start";
      case "CENTER":
        return "center";
      case "BOTTOM":
        return "flex-end";
    }
  };

  return (
    <div
      css={css`
        @media (max-width: 768px) {
          width: 100vw;
          height: 100vh;
        }
        width: 768px;
        min-height: 100vh;
        overflow: auto;

        background-color: ${dark1MainBg};

        display: flex;
        flex-direction: column;
        justify-content: ${justifyContent(position)};
      `}
    >
      {children}
    </div>
  );
};
