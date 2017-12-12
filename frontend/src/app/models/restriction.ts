export class Restriction {
    constructor(
    public operator: string = "=",
    public type: RestrictionType = RestrictionType.END,
    public value: string = "") {}

    static typeString(type: RestrictionType): string {
        switch (type) {
            case RestrictionType.END:
            default:
                return "Ends with:"
        }
    }

}

export enum RestrictionType {
    END = "END"
}
