"use client";
import { useEffect, useState } from "react";
import { FlickeringTrail } from "./FlickeringTrail";

interface Props {
  command: string;
}

export function CommandTypeIn(props: Props) {
  const [writtenLength, setWrittenLength] = useState(0);
  const command = props.command;

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
}
