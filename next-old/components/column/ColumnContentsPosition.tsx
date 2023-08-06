import { ReactNode } from "react";
import { ColumnVerticalPosition } from "../../libs/gql/graphql";
import { dark1MainBg, dark3SubBg } from "../../libs/colorTheme";
import { css } from "@emotion/react";

interface ColumnContentsPositionProps {
  position: ColumnVerticalPosition;
  children: ReactNode;
}

export const ColumnContentsPosition = ({
  position,
  children,
}: ColumnContentsPositionProps) => {
  const justifyContent = (p: ColumnVerticalPosition): string => {
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
        /* min-height: 100vh; */
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
