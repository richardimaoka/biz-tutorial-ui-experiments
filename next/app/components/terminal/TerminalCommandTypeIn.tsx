"use client";
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
    <code>
      {command?.substring(0, writtenLength)}
      {writtenLength >= command?.length && <FlickeringTrail />}
    </code>
  );
};
