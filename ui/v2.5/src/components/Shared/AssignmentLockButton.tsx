import React from "react";
import cx from "classnames";
import { Button } from "react-bootstrap";
import { faLock, faLockOpen } from "@fortawesome/free-solid-svg-icons";
import { Icon } from "src/components/Shared/Icon";

interface IAssignmentLockButtonProps {
  locked: boolean;
  label: string;
  onToggle: () => void;
  className?: string;
}

export const AssignmentLockButton: React.FC<IAssignmentLockButtonProps> = ({
  locked,
  label,
  onToggle,
  className,
}) => (
  <Button
    type="button"
    variant="secondary"
    title={label}
    aria-label={label}
    aria-pressed={locked}
    className={cx(
      "minimal",
      "assignment-lock-button",
      locked ? "locked" : "unlocked",
      className
    )}
    onClick={onToggle}
  >
    <Icon icon={locked ? faLock : faLockOpen} />
  </Button>
);
