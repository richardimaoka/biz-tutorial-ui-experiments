"use client";
import { source_code_pro } from "../fonts/fonts";
import { useEffect, useState } from "react";
import { FlickeringTrail } from "./FlickeringTrail";

interface TerminalCommandTypeInProps {
  command: string;
}

export const TerminalCommandTypeIn = ({
  command,
}: TerminalCommandTypeInProps) => {
  const [writtenLength, setWrittenLength] = useState(0);

  useEffect(() => {
    if (command && writtenLength < command.length) {
      const incrementStep = command.length / 10;
      const nextLength = Math.min(
        writtenLength + incrementStep,
        command.length
      );
      setTimeout(() => {
        setWrittenLength(nextLength);
      }, 20);
    }
  });

  return (
    <code
      // needs to specify font here, as <code> has its own font
      className={source_code_pro.className}
    >
      {command?.substring(0, writtenLength)}
      {writtenLength >= command?.length && <FlickeringTrail />}
    </code>
  );
};
