// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Immutable;

namespace Pulumi.Mongodbatlas
{
    public static class Config
    {
        [System.Diagnostics.CodeAnalysis.SuppressMessage("Microsoft.Design", "IDE1006", Justification = 
        "Double underscore prefix used to avoid conflicts with variable names.")]
        private sealed class __Value<T>
        {
            private readonly Func<T> _getter;
            private T _value = default!;
            private bool _set;

            public __Value(Func<T> getter)
            {
                _getter = getter;
            }

            public T Get() => _set ? _value : _getter();

            public void Set(T value)
            {
                _value = value;
                _set = true;
            }
        }

        private static readonly Pulumi.Config __config = new Pulumi.Config("mongodbatlas");

        private static readonly __Value<string?> _baseUrl = new __Value<string?>(() => __config.Get("baseUrl"));
        /// <summary>
        /// MongoDB Atlas Base URL
        /// </summary>
        public static string? BaseUrl
        {
            get => _baseUrl.Get();
            set => _baseUrl.Set(value);
        }

        private static readonly __Value<string?> _privateKey = new __Value<string?>(() => __config.Get("privateKey"));
        /// <summary>
        /// MongoDB Atlas Programmatic Private Key
        /// </summary>
        public static string? PrivateKey
        {
            get => _privateKey.Get();
            set => _privateKey.Set(value);
        }

        private static readonly __Value<string?> _publicKey = new __Value<string?>(() => __config.Get("publicKey"));
        /// <summary>
        /// MongoDB Atlas Programmatic Public Key
        /// </summary>
        public static string? PublicKey
        {
            get => _publicKey.Get();
            set => _publicKey.Set(value);
        }

        private static readonly __Value<string?> _realmBaseUrl = new __Value<string?>(() => __config.Get("realmBaseUrl"));
        /// <summary>
        /// MongoDB Realm Base URL
        /// </summary>
        public static string? RealmBaseUrl
        {
            get => _realmBaseUrl.Get();
            set => _realmBaseUrl.Set(value);
        }

    }
}
