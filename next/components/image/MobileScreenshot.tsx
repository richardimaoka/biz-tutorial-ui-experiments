import { css } from "@emotion/react";
import Image from "next/image";

interface MobileScreenshotProps {
  width: number;
  height: number;
  src: string;
}

export const MobileScreenshot = ({
  width,
  height,
  src,
}: MobileScreenshotProps) => (
  <Image
    css={css`
      display: block;
      margin: 0 auto;
      @media (max-width: 768px) {
        width: 100%;
        height: auto;
      }
    `}
    src={src}
    width={width}
    height={height}
    alt="screenshot on mobile"
  />
);
