// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.mongodbatlas.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetMaintenanceWindowResult {
    /**
     * @return Flag that indicates whether you want to defer all maintenance windows one week they would be triggered.
     * For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/maintenance-windows/)
     * 
     */
    private Boolean autoDeferOnceEnabled;
    /**
     * @return Day of the week when you would like the maintenance window to start as a 1-based integer: S=1, M=2, T=3, W=4, T=5, F=6, S=7.
     * 
     */
    private Integer dayOfWeek;
    /**
     * @return Hour of the day when you would like the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0, noon is 12  (Time zone is UTC).
     * 
     */
    private Integer hourOfDay;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Number of times the current maintenance event for this project has been deferred, you can set a maximum of 2 deferrals.
     * 
     */
    private Integer numberOfDeferrals;
    private String projectId;
    /**
     * @return Flag indicating whether project maintenance has been directed to start immediately. If you request that maintenance begin immediately, this field returns true from the time the request was made until the time the maintenance event completes.
     * 
     */
    private Boolean startAsap;

    private GetMaintenanceWindowResult() {}
    /**
     * @return Flag that indicates whether you want to defer all maintenance windows one week they would be triggered.
     * For more information see: [MongoDB Atlas API Reference.](https://docs.atlas.mongodb.com/reference/api/maintenance-windows/)
     * 
     */
    public Boolean autoDeferOnceEnabled() {
        return this.autoDeferOnceEnabled;
    }
    /**
     * @return Day of the week when you would like the maintenance window to start as a 1-based integer: S=1, M=2, T=3, W=4, T=5, F=6, S=7.
     * 
     */
    public Integer dayOfWeek() {
        return this.dayOfWeek;
    }
    /**
     * @return Hour of the day when you would like the maintenance window to start. This parameter uses the 24-hour clock, where midnight is 0, noon is 12  (Time zone is UTC).
     * 
     */
    public Integer hourOfDay() {
        return this.hourOfDay;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Number of times the current maintenance event for this project has been deferred, you can set a maximum of 2 deferrals.
     * 
     */
    public Integer numberOfDeferrals() {
        return this.numberOfDeferrals;
    }
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return Flag indicating whether project maintenance has been directed to start immediately. If you request that maintenance begin immediately, this field returns true from the time the request was made until the time the maintenance event completes.
     * 
     */
    public Boolean startAsap() {
        return this.startAsap;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetMaintenanceWindowResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean autoDeferOnceEnabled;
        private Integer dayOfWeek;
        private Integer hourOfDay;
        private String id;
        private Integer numberOfDeferrals;
        private String projectId;
        private Boolean startAsap;
        public Builder() {}
        public Builder(GetMaintenanceWindowResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.autoDeferOnceEnabled = defaults.autoDeferOnceEnabled;
    	      this.dayOfWeek = defaults.dayOfWeek;
    	      this.hourOfDay = defaults.hourOfDay;
    	      this.id = defaults.id;
    	      this.numberOfDeferrals = defaults.numberOfDeferrals;
    	      this.projectId = defaults.projectId;
    	      this.startAsap = defaults.startAsap;
        }

        @CustomType.Setter
        public Builder autoDeferOnceEnabled(Boolean autoDeferOnceEnabled) {
            this.autoDeferOnceEnabled = Objects.requireNonNull(autoDeferOnceEnabled);
            return this;
        }
        @CustomType.Setter
        public Builder dayOfWeek(Integer dayOfWeek) {
            this.dayOfWeek = Objects.requireNonNull(dayOfWeek);
            return this;
        }
        @CustomType.Setter
        public Builder hourOfDay(Integer hourOfDay) {
            this.hourOfDay = Objects.requireNonNull(hourOfDay);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder numberOfDeferrals(Integer numberOfDeferrals) {
            this.numberOfDeferrals = Objects.requireNonNull(numberOfDeferrals);
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        @CustomType.Setter
        public Builder startAsap(Boolean startAsap) {
            this.startAsap = Objects.requireNonNull(startAsap);
            return this;
        }
        public GetMaintenanceWindowResult build() {
            final var o = new GetMaintenanceWindowResult();
            o.autoDeferOnceEnabled = autoDeferOnceEnabled;
            o.dayOfWeek = dayOfWeek;
            o.hourOfDay = hourOfDay;
            o.id = id;
            o.numberOfDeferrals = numberOfDeferrals;
            o.projectId = projectId;
            o.startAsap = startAsap;
            return o;
        }
    }
}
