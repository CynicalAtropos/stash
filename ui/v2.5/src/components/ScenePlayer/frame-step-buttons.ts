import videojs, { VideoJsPlayer } from "video.js";

interface IFrameStepButtonsOptions {
  frameRate?: number;
}

interface FrameStepButtonOptions extends videojs.ComponentOptions {
  direction: "forward" | "back";
  parent: FrameStepButtonsPlugin;
}

function isValidFrameRate(frameRate: number | undefined): frameRate is number {
  return frameRate !== undefined && Number.isFinite(frameRate) && frameRate > 0;
}

class FrameStepButtonsPlugin extends videojs.getPlugin("plugin") {
  private frameRate?: number;
  private backButton: FrameStepButton;
  private forwardButton: FrameStepButton;

  constructor(player: VideoJsPlayer, options?: IFrameStepButtonsOptions) {
    super(player, options);

    this.frameRate = options?.frameRate;
    this.backButton = new FrameStepButton(player, {
      direction: "back",
      parent: this,
    });
    this.forwardButton = new FrameStepButton(player, {
      direction: "forward",
      parent: this,
    });

    player.ready(() => {
      this.ready();
    });
  }

  private ready() {
    const { controlBar } = this.player;

    this.player.addClass("vjs-frame-step-buttons");
    controlBar.addChild(this.backButton);
    controlBar.addChild(this.forwardButton);

    window.setTimeout(() => this.placeButtons(), 0);
    this.updateButtonVisibility();
  }

  private placeButtons() {
    const { controlBar } = this.player;
    const playToggle = controlBar.getChild("playToggle");

    if (playToggle) {
      const playToggleEl = playToggle.el();
      controlBar.el().insertBefore(this.backButton.el(), playToggleEl);
      controlBar
        .el()
        .insertBefore(this.forwardButton.el(), playToggleEl.nextSibling);
    }
  }

  public setFrameRate(frameRate: number | undefined) {
    this.frameRate = frameRate;
    this.updateButtonVisibility();
  }

  public step(direction: "forward" | "back") {
    const frameRate = this.frameRate;
    if (!isValidFrameRate(frameRate)) return;

    const currentTime = this.player.currentTime();
    const duration = this.player.duration();
    const frameDuration = 1 / frameRate;
    const delta = direction === "forward" ? frameDuration : -frameDuration;
    let nextTime = Math.max(0, currentTime + delta);

    if (Number.isFinite(duration)) {
      nextTime = Math.min(duration, nextTime);
    }

    this.player.pause();
    this.player.currentTime(nextTime);
  }

  private updateButtonVisibility() {
    const visible = isValidFrameRate(this.frameRate);
    this.backButton.setVisible(visible);
    this.forwardButton.setVisible(visible);

    if (visible) {
      this.player.addClass("vjs-frame-step-buttons-enabled");
    } else {
      this.player.removeClass("vjs-frame-step-buttons-enabled");
    }
  }
}

class FrameStepButton extends videojs.getComponent("Button") {
  private parentPlugin: FrameStepButtonsPlugin;
  private direction: "forward" | "back";

  constructor(player: VideoJsPlayer, options: FrameStepButtonOptions) {
    super(player, options);
    this.parentPlugin = options.parent;
    this.direction = options.direction;

    if (options.direction === "forward") {
      this.controlText(this.localize("Next frame"));
    } else {
      this.controlText(this.localize("Previous frame"));
    }

    const icon = document.createElement("span");
    icon.className = "vjs-frame-step-icon";
    icon.setAttribute("aria-hidden", "true");
    icon.textContent = options.direction === "forward" ? "|>" : "<|";
    this.el().appendChild(icon);
  }

  buildCSSClass() {
    return `vjs-frame-step-button vjs-frame-step-${this.direction} ${super.buildCSSClass()}`;
  }

  handleClick(event: Event) {
    event.stopPropagation();
    this.parentPlugin.step(this.direction);
  }

  public setVisible(visible: boolean) {
    if (visible) {
      this.removeClass("vjs-hidden");
    } else {
      this.addClass("vjs-hidden");
    }
  }
}

videojs.registerComponent("FrameStepButton", FrameStepButton);
videojs.registerPlugin("frameStepButtons", FrameStepButtonsPlugin);

declare module "video.js" {
  interface VideoJsPlayer {
    frameStepButtons: () => FrameStepButtonsPlugin;
  }
  interface VideoJsPlayerPluginOptions {
    frameStepButtons?: IFrameStepButtonsOptions;
  }
}

export default FrameStepButtonsPlugin;
